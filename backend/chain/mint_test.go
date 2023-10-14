package chain

import (
	"github.com/paulgoleary/local-luv-proto/crypto"
	"github.com/stretchr/testify/require"
	"github.com/umbracle/ethgo"
	"github.com/umbracle/ethgo/contract"
	"github.com/umbracle/ethgo/jsonrpc"
	"math/big"
	"os"
	"testing"
)

var billyzWallet = ethgo.HexToAddress("0xEF17dc60E4D58Fd24D3b6FCDF07e3C5029018863")

func noTestSFLUVMint(t *testing.T) {
	ec, err := jsonrpc.NewClient(os.Getenv("CHAIN_URL"))
	require.NoError(t, err)

	sk, err := crypto.SKFromHex(os.Getenv("CHAIN_SK"))
	require.NoError(t, err)
	k := &EcdsaKey{SK: sk}

	if false {
		ep, err := LoadContract(ec, "ERC20.sol/ERC20", k, USDCPolygonMainnet)
		require.NoError(t, err)

		// function approve(address spender, uint256 amount) public virtual override returns (bool)
		err = TxnDoWait(ep.Txn("approve", SFLUVPolygonMainnetV1_1, big.NewInt(1*1_000_000)))
		require.NoError(t, err)

		// function allowance(address owner, address spender) external view returns (uint256);
		resp, err := ep.Call("allowance", ethgo.Latest, k.Address(), SFLUVPolygonMainnetV1_1)
		require.NoError(t, err)
		_ = resp
	}

	ep, err := LoadContract(ec, "SFLUVv1.sol/SFLUVv1", k, SFLUVPolygonMainnetV1_1)
	require.NoError(t, err)

	if true {
		resp, err := ep.Call("MINTER_ROLE", ethgo.Latest)
		require.NoError(t, err)
		minterRole, ok := resp["0"].([32]byte)
		require.True(t, ok)

		resp, err = ep.Call("hasRole", ethgo.Latest, minterRole, k.Address())
		require.NoError(t, err)
		hasRole, ok := resp["0"].(bool)
		require.True(t, ok)

		if !hasRole {
			// testLUVCoin.grantRole(testLUVCoin.MINTER_ROLE(), testLUVCoin.owner());
			err = TxnDoWait(ep.Txn("grantRole", minterRole, k.Address()))
			require.NoError(t, err)
		}

		// err = TxnDoWait(ep.Txn("depositFor", k.Address(), big.NewInt(100*1_000_000)))
		err = TxnDoWait(ep.Txn("depositFor", billyzWallet, big.NewInt(1_000)))
		require.NoError(t, err)

		resp, err = ep.Call("balanceOf", ethgo.Latest, k.Address())
		require.NoError(t, err)
		_ = resp
	}

	if false {
		// function withdrawTo(address account, uint256 amount) public virtual returns (bool) {
		err = TxnDoWait(ep.Txn("withdrawTo", k.Address(), big.NewInt(100*1_000_000)))
		require.NoError(t, err)

		resp, err := ep.Call("balanceOf", ethgo.Latest, k.Address())
		require.NoError(t, err)
		_ = resp
	}

}

func noTestMockCoinMint(t *testing.T) {

	ec, err := jsonrpc.NewClient(os.Getenv("CHAIN_URL"))
	require.NoError(t, err)

	sk, err := crypto.SKFromHex(os.Getenv("CHAIN_SK"))
	require.NoError(t, err)
	k := &EcdsaKey{SK: sk}

	ep, err := LoadContract(ec, "MockCoin.sol/MockCoin", k, MockMumbaiAddr)
	require.NoError(t, err)

	walletAddr := ethgo.HexToAddress("0x054dF6203225bB58d9243eBf9DAd55608a436042")

	//res, err := ep.Call("balanceOf", ethgo.Latest, walletAddr)
	//require.NoError(t, err)
	//checkBalance, ok := res["0"].(*big.Int)
	//require.True(t, ok)
	//require.Equal(t, 0, checkBalance.Cmp(big.NewInt(0)))

	tx, err := ep.Txn("mint", walletAddr, ethgo.Ether(100))
	require.NoError(t, err)
	tx.WithOpts(&contract.TxnOpts{GasLimit: 100_000})

	err = TxnDoWait(tx, nil)
	require.NoError(t, err)

	res, err := ep.Call("balanceOf", ethgo.Latest, walletAddr)
	require.NoError(t, err)
	checkBalance, ok := res["0"].(*big.Int)
	require.True(t, ok)
	require.Equal(t, 0, checkBalance.Cmp(ethgo.Ether(200)))
}
