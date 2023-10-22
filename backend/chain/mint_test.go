package chain

import (
	"github.com/stretchr/testify/require"
	"github.com/umbracle/ethgo"
	"math/big"
	"os"
	"testing"
)

var billyzWallet = ethgo.HexToAddress("0xEF17dc60E4D58Fd24D3b6FCDF07e3C5029018863")

func noTestSFLUVMint(t *testing.T) {

	mh, err := makeMintHelper(os.Getenv("CHAIN_URL"), os.Getenv("CHAIN_SK"), LuvMumbaiAddr, MockMumbaiAddr)
	require.NoError(t, err)

	err = mh.mintLuv(LuvGovMumbaiAddr, ethgo.Ether(1000))
	require.NoError(t, err)

	resp, err := mh.luvCoin.Call("balanceOf", ethgo.Latest, LuvGovMumbaiAddr)
	require.NoError(t, err)
	checkAmt, ok := resp["0"].(*big.Int)
	require.True(t, ok)
	require.Equal(t, 0, checkAmt.Cmp(ethgo.Ether(1000)))
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
