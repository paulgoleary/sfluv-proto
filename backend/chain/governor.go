package chain

import (
	"crypto/ecdsa"
	"github.com/paulgoleary/local-luv-proto/crypto"
	"github.com/umbracle/ethgo"
	"github.com/umbracle/ethgo/contract"
	"github.com/umbracle/ethgo/jsonrpc"
)

func deployGovernance(rpcUrl, skHex string) (votesAddr, govAddr ethgo.Address, err error) {

	var ec *jsonrpc.Client
	if ec, err = jsonrpc.NewClient(rpcUrl); err != nil {
		return
	}

	var sk *ecdsa.PrivateKey
	if sk, err = crypto.SKFromHex(skHex); err != nil {
		return
	}

	k := &EcdsaKey{SK: sk}

	if _, votesAddr, err = deployContract(ec, "SFLUVVotes.sol/SFLUVVotesV1", k, nil); err != nil {
		return
	}

	if _, govAddr, err = deployContract(ec, "SFLUVGovernor.sol/SFLUVGovernorV0", k, []interface{}{votesAddr}); err != nil {
		return
	}

	return
}

type govHelper struct {
	gov   *contract.Contract
	votes *contract.Contract

	k *EcdsaKey
}

func makeGovHelper(rpcUrl, skHex string, govAddr, votesAddr ethgo.Address) (gh *govHelper, err error) {

	var ec *jsonrpc.Client
	if ec, err = jsonrpc.NewClient(rpcUrl); err != nil {
		return
	}

	var sk *ecdsa.PrivateKey
	if sk, err = crypto.SKFromHex(skHex); err != nil {
		return
	}

	gh = &govHelper{k: &EcdsaKey{SK: sk}}

	if gh.gov, err = LoadContract(ec, "SFLUVGovernor.sol/SFLUVGovernorV0", gh.k, govAddr); err != nil {
		return
	}

	if gh.votes, err = LoadContract(ec, "SFLUVVotes.sol/SFLUVVotesV1", gh.k, votesAddr); err != nil {
		return
	}

	return
}

func (gh *govHelper) mintVote(toAddr ethgo.Address) (err error) {
	if err = checkMinterRole(gh.votes, gh.k.Address()); err != nil {
		return
	}

	if err = TxnDoWait(gh.votes.Txn("mint", toAddr)); err != nil {
		return
	}

	return
}
