package erc4337

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/paulgoleary/local-luv-proto/chain"
	"github.com/paulgoleary/local-luv-proto/crypto"
	"github.com/stretchr/testify/require"
	"github.com/umbracle/ethgo"
	"github.com/umbracle/ethgo/jsonrpc"
	"github.com/umbracle/ethgo/jsonrpc/codec"
	"os"
	"testing"
)

func TestInitCode(t *testing.T) {

	sk, err := crypto.SKFromHex(os.Getenv("CHAIN_SK"))
	require.NoError(t, err)
	k := &chain.EcdsaKey{SK: sk}

	testInitCode, err := MakeDefaultInitCode(k.Address())
	require.NoError(t, err)
	require.Equal(t, 88, len(testInitCode))

	// account creation gas usage: 288,526 per Mumbai
}

func TestGetSenderAddress(t *testing.T) {

	ec, err := jsonrpc.NewClient(os.Getenv("CHAIN_URL"))
	require.NoError(t, err)

	sk, err := crypto.SKFromHex(os.Getenv("CHAIN_SK"))
	require.NoError(t, err)
	k := &chain.EcdsaKey{SK: sk}

	require.Equal(t, "0x32A629dE3fb4549EB2B204d37eb9C8CFb0b9AdCf", k.Address().String())

	ep, err := chain.LoadContract(ec, "IEntryPoint.sol/IEntryPoint", k, DefaultEntryPoint)
	require.NoError(t, err)

	initCode, err := MakeDefaultInitCode(k.Address())
	require.NoError(t, err)

	_, err = ep.Call("getSenderAddress", ethgo.Latest, initCode)
	// this method is expected to revert
	eo, ok := err.(*codec.ErrorObject)
	require.True(t, ok)
	eod, ok := eo.Data.(string)
	require.True(t, ok)

	eodBytes, err := hexutil.Decode(eod)
	require.NoError(t, err)

	require.Equal(t, 36, len(eodBytes))
	senderAddrBytes := eodBytes[16:]
	require.Equal(t, ethgo.HexToAddress("0x054dF6203225bB58d9243eBf9DAd55608a436042"), ethgo.BytesToAddress(senderAddrBytes))

}

func TestInitABI(t *testing.T) {

	abiAccountFactory, err := chain.LoadABI("SimpleAccountFactory.sol/SimpleAccountFactory")
	require.NoError(t, err)

	checkCreateMethod := abiAccountFactory.GetMethod("createAccount")

	require.Equal(t, checkCreateMethod.Sig(), createAccountMethod.Sig())
}
