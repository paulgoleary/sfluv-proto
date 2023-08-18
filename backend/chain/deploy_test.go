package chain

import (
	"github.com/paulgoleary/local-luv-proto/crypto"
	"github.com/stretchr/testify/require"
	"github.com/umbracle/ethgo"
	"github.com/umbracle/ethgo/jsonrpc"
	"os"
	"testing"
)

var USDCPolygonMainnet = ethgo.HexToAddress("0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174")

func TestDeploy(t *testing.T) {
	ec, err := jsonrpc.NewClient(os.Getenv("CHAIN_URL"))
	require.NoError(t, err)

	sk, err := crypto.SKFromHex(os.Getenv("CHAIN_SK"))
	require.NoError(t, err)
	k := &EcdsaKey{SK: sk}

	//_, mockAddr, err := deployContract(ec, "MockCoin.sol/MockCoin", k, nil)
	//require.NoError(t, err)

	_, luvAddr, err := deployContract(ec, "SFLUVv1.sol/SFLUVv1", k, []interface{}{USDCPolygonMainnet})
	require.NoError(t, err)

	println(luvAddr.String())
}
