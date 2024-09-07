package chain

import (
	"github.com/paulgoleary/local-luv-proto/util"
	"github.com/stretchr/testify/require"
	"github.com/umbracle/ethgo"
	"math/big"
	"os"
	"testing"
)

var billyzWallet = ethgo.HexToAddress("0xEF17dc60E4D58Fd24D3b6FCDF07e3C5029018863")
var paulzWallet = ethgo.HexToAddress("0x05c0486C000C7a09F497ec1680e2aC540133342E")

var sfLUVNewAdminAddr = ethgo.HexToAddress("0xF1c338b02c589d057D23Db7f379f3499D228b4A5")

func TestSFLUVMint(t *testing.T) {

	mh, err := makeMintHelper(os.Getenv("CHAIN_URL"), os.Getenv("CHAIN_SK"), SFLUVPolygonMainnetV1_1, USDCPolygonMainnet)
	require.NoError(t, err)

	mintAmt := util.Mega(500)

	err = mh.mintLuv(sfLUVNewAdminAddr, mintAmt)
	require.NoError(t, err)

	resp, err := mh.luvCoin.Call("balanceOf", ethgo.Latest, sfLUVNewAdminAddr)
	require.NoError(t, err)
	checkAmt, ok := resp["0"].(*big.Int)
	require.True(t, ok)
	require.True(t, checkAmt.Cmp(mintAmt) >= 0)
}

func noTestMockCoinMint(t *testing.T) {
	mh, err := makeMintHelper(os.Getenv("CHAIN_URL"), os.Getenv("CHAIN_SK"), LuvMumbaiAddr, MockMumbaiAddr)
	require.NoError(t, err)

	err = mh.mintBase(mh.k.Address(), ethgo.Ether(1000))
	require.NoError(t, err)

	res, err := mh.baseCoin.Call("balanceOf", ethgo.Latest, mh.k.Address())
	require.NoError(t, err)
	checkBalance, ok := res["0"].(*big.Int)
	require.True(t, ok)
	require.Equal(t, 0, checkBalance.Cmp(ethgo.Ether(1000)))
}
