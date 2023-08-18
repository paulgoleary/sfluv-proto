package erc4337

import (
	"fmt"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/paulgoleary/local-luv-proto/chain"
	"github.com/paulgoleary/local-luv-proto/crypto"
	"github.com/stretchr/testify/require"
	"github.com/umbracle/ethgo"
	"github.com/umbracle/ethgo/jsonrpc"
	"math/big"
	"os"
	"testing"
)

func TestSendUserOp(t *testing.T) {

	// https://api.stackup.sh/v1/node/5a5fed50fc4f446099f4628efa9c34b64339b2e56376c70da18b13c9ac89bb16
	nodeRpc, err := rpc.Dial("https://api.stackup.sh/v1/node/5a5fed50fc4f446099f4628efa9c34b64339b2e56376c70da18b13c9ac89bb16")
	require.NoError(t, err)

	pmRpc, err := rpc.Dial("https://api.stackup.sh/v1/paymaster/5a5fed50fc4f446099f4628efa9c34b64339b2e56376c70da18b13c9ac89bb16")
	require.NoError(t, err)

	sk, err := crypto.SKFromHex(os.Getenv("CHAIN_SK"))
	require.NoError(t, err)
	k := &chain.EcdsaKey{SK: sk}

	ec, err := jsonrpc.NewClient(os.Getenv("CHAIN_URL"))
	require.NoError(t, err)

	ep, err := chain.LoadContract(ec, "IEntryPoint.sol/IEntryPoint", k, DefaultEntryPoint)
	require.NoError(t, err)

	res, err := ep.Call("getNonce", ethgo.Latest,
		ethgo.HexToAddress("0x054dF6203225bB58d9243eBf9DAd55608a436042"), big.NewInt(0))
	require.NoError(t, err)
	nonce, ok := res["nonce"].(*big.Int)
	require.True(t, ok)

	op, err := UserOpApprove(
		nonce,
		ethgo.HexToAddress("0x32A629dE3fb4549EB2B204d37eb9C8CFb0b9AdCf"),
		ethgo.HexToAddress("0x054dF6203225bB58d9243eBf9DAd55608a436042"),
		chain.MockMumbaiAddr,
		chain.LuvMumbaiAddr, // 'spender' should be sfluv contract
		ethgo.Ether(100))
	require.NoError(t, err)

	op, err = UserOpSeal(op, big.NewInt(80001), k)
	require.NoError(t, err)

	opMap, _ := op.ToMap()

	var pmResp map[string]any
	err = pmRpc.Call(&pmResp, "pm_sponsorUserOperation", opMap, DefaultEntryPoint.String(), map[string]string{"type": "payg"})
	require.NoError(t, err)

	op, err = UserOpSeal(op, big.NewInt(80001), k)
	require.NoError(t, err)

	opMap, _ = op.ToMap()

	var reply string
	err = nodeRpc.Call(&reply, "eth_sendUserOperation", opMap, DefaultEntryPoint.String())
	require.NoError(t, err)
	fmt.Printf("Response: %v", reply)

}
