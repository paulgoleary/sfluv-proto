package erc4337

import (
	"github.com/paulgoleary/local-luv-proto/chain"
	"github.com/stretchr/testify/require"
	"github.com/umbracle/ethgo"
	"github.com/umbracle/ethgo/jsonrpc"
	"os"
	"testing"
)

func TestCWAccount(t *testing.T) {

	ec, err := jsonrpc.NewClient(os.Getenv("CHAIN_URL"))
	require.NoError(t, err)

	pgoWalletAddr := ethgo.HexToAddress("0x05c0486C000C7a09F497ec1680e2aC540133342E")
	k := &chain.AddressKey{Addr: pgoWalletAddr}

	ep, err := chain.LoadContract(ec, "Account.sol/Account", k, pgoWalletAddr)
	require.NoError(t, err)

	res, err := ep.Call("tokenEntryPoint", ethgo.Latest)
	tepAddr, ok := res["0"].(ethgo.Address)
	require.True(t, ok)
	println(tepAddr.String())
}
