package chain

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/paulgoleary/local-luv-proto/crypto"
	"github.com/umbracle/ethgo"
	"github.com/umbracle/ethgo/contract"
	"github.com/umbracle/ethgo/jsonrpc"
	"math/big"
)

type mintHelper struct {
	baseCoin *contract.Contract
	luvCoin  *contract.Contract

	luvAddr ethgo.Address
	k       *EcdsaKey
}

func makeMintHelper(rpcUrl, skHex string, luvAddr, baseAddr ethgo.Address) (mh *mintHelper, err error) {

	var ec *jsonrpc.Client
	if ec, err = jsonrpc.NewClient(rpcUrl); err != nil {
		return
	}

	var sk *ecdsa.PrivateKey
	if sk, err = crypto.SKFromHex(skHex); err != nil {
		return
	}

	mh = &mintHelper{k: &EcdsaKey{SK: sk}, luvAddr: luvAddr}

	if mh.luvCoin, err = LoadContract(ec, "SFLUVv1.sol/SFLUVv1", mh.k, luvAddr); err != nil {
		return
	}

	// MockCoin is ERC20 but also has the 'mint' method in case we need to mint base coins
	if mh.baseCoin, err = LoadContract(ec, "MockCoin.sol/MockCoin", mh.k, baseAddr); err != nil {
		return
	}

	return
}

// check that the provided key has the MINTER role and - if not - attempt to add it
// note that the configured key has to be an ADMIN of the MINTER role
func checkMinterRole(c *contract.Contract, toAddr ethgo.Address) error {
	resp, err := c.Call("MINTER_ROLE", ethgo.Latest)
	if err != nil {
		return err
	}
	minterRole, ok := resp["0"].([32]byte)
	if !ok {
		return fmt.Errorf("should not happen - target contract does not support MINTER_ROLE")
	}

	resp, err = c.Call("hasRole", ethgo.Latest, minterRole, toAddr)
	if err != nil {
		return err
	}
	hasRole, ok := resp["0"].(bool)

	if !hasRole {
		if err = TxnDoWait(c.Txn("grantRole", minterRole, toAddr)); err != nil {
			return err
		}
	}

	return nil
}

// this only works if we're allowed to mint against the base coin contract - e.g. it's the mock coin
func (mh *mintHelper) mintBase(toAddr ethgo.Address, amt *big.Int) (err error) {

	var tx contract.Txn
	if tx, err = mh.baseCoin.Txn("mint", toAddr, amt); err != nil {
		return
	}
	tx.WithOpts(&contract.TxnOpts{GasLimit: 100_000})

	if err = TxnDoWait(tx, nil); err != nil {
		return
	}

	return
}

// assumes provided key owns sufficient amount in base coin
func (mh *mintHelper) mintLuv(toAddr ethgo.Address, amt *big.Int) (err error) {

	if err = checkMinterRole(mh.luvCoin, mh.k.Address()); err != nil {
		return
	}

	// function allowance(address owner, address spender) external view returns (uint256);
	var resp map[string]any
	if resp, err = mh.baseCoin.Call("allowance", ethgo.Latest, mh.k.Address(), mh.luvAddr); err != nil {
		return
	}
	allowance, ok := resp["0"].(*big.Int)
	if !ok {
		err = fmt.Errorf("should not happen - 'allowance' method did not return proper result")
		return
	}
	if allowance.Cmp(amt) < 0 {
		// current allowance is not sufficient
		allowance.Sub(amt, allowance)
		// function approve(address spender, uint256 amount) public virtual override returns (bool)
		if err = TxnDoWait(mh.baseCoin.Txn("approve", mh.luvAddr, allowance)); err != nil {
			return
		}
	}

	if err = TxnDoWait(mh.luvCoin.Txn("depositFor", toAddr, allowance)); err != nil {
		return
	}

	return
}
