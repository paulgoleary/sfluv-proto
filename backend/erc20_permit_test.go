package local_luv_proto

import (
	"github.com/paulgoleary/local-luv-proto/crypto"
	"github.com/stretchr/testify/require"
	"github.com/umbracle/ethgo"
	"github.com/umbracle/ethgo/jsonrpc"
	"os"
	"testing"
)

// TODO: these are currently older versions... in particular, no mint function on mock...
var mockMumbaiAddr = ethgo.HexToAddress("0x834F9b26Cc7C806c6F9f31697C4B1C20A1bB83b6")
var luvMumbaiAddr = ethgo.HexToAddress("0x281E6f63E932392135D7A3847C103CF0F358865F")

func noTestDeploy(t *testing.T) {
	ec, err := jsonrpc.NewClient(os.Getenv("NETWORK_URL")) // currently mumbai
	require.NoError(t, err)

	sk, err := crypto.SKFromHex(os.Getenv("NETWORK_SK"))
	require.NoError(t, err)
	k := &ecdsaKey{k: sk}

	_, mockAddr, err := deployArtifact(ec, "MockCoin.sol/MockCoin", k, nil)
	require.NoError(t, err)

	_, luvAddr, err := deployArtifact(ec, "SFLUVv1.sol/SFLUVv1", k, []interface{}{mockAddr})
	require.NoError(t, err)

	println(mockAddr.String())
	println(luvAddr.String())
}
