package chain

import (
	"github.com/paulgoleary/local-luv-proto/crypto"
	"github.com/stretchr/testify/require"
	"github.com/umbracle/ethgo/jsonrpc"
	"os"
	"testing"
)

func noTestDeploy(t *testing.T) {
	ec, err := jsonrpc.NewClient(os.Getenv("NETWORK_URL")) // currently mumbai
	require.NoError(t, err)

	sk, err := crypto.SKFromHex(os.Getenv("NETWORK_SK"))
	require.NoError(t, err)
	k := &EcdsaKey{SK: sk}

	//_, mockAddr, err := deployContract(ec, "MockCoin.sol/MockCoin", SK, nil)
	//require.NoError(t, err)

	_, luvAddr, err := deployContract(ec, "SFLUVv1.sol/SFLUVv1", k, []interface{}{MockMumbaiAddr})
	require.NoError(t, err)

	println(luvAddr.String())
}
