package chain

import (
	"fmt"
	"github.com/paulgoleary/local-luv-proto/crypto"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestDAOMumbai(t *testing.T) {

	gh, err := makeGovHelper(os.Getenv("CHAIN_URL"), os.Getenv("CHAIN_SK"), LuvGovMumbaiAddr, LuvVotesMumbaiAddr)
	require.NoError(t, err)

	for i := 0; i < 3; i++ {
		sk, _ := crypto.RandSK()
		ec := EcdsaKey{SK: sk}

		err = gh.mintVote(ec.Address())
		require.NoError(t, err)

		fmt.Printf("address '%v', sk '%v'\n", ec.Address().String(), ec.SKHex())
	}
}
