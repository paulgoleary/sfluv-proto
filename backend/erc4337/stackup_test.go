package erc4337

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/paulgoleary/local-luv-proto/chain"
	"github.com/paulgoleary/local-luv-proto/crypto"
	"github.com/stackup-wallet/stackup-bundler/pkg/entrypoint/reverts"
	"github.com/stackup-wallet/stackup-bundler/pkg/userop"
	"github.com/stretchr/testify/require"
	"github.com/umbracle/ethgo/jsonrpc"
	"github.com/umbracle/ethgo/jsonrpc/codec"
	"math/big"
	"os"
	"strings"
	"testing"
)

func TestCheckUserOp(t *testing.T) {
	nodeRpc, err := rpc.Dial(os.Getenv("SU_NODE_URL"))
	require.NoError(t, err)

	opHash := "0x1410c65c614062a0e3885caf8750d35b3a25d2e0c4506a572aac2f11c6052ac0"

	var reply map[string]any
	err = nodeRpc.Call(&reply, "eth_getUserOperationReceipt", opHash)
	require.NoError(t, err)

	fmt.Printf("Response: %v", reply)
}

func TestSendUserOp(t *testing.T) {

	nodeRpc, err := rpc.Dial(os.Getenv("SU_NODE_URL"))
	require.NoError(t, err)

	pmRpc, err := rpc.Dial(os.Getenv("SU_PM_URL"))
	require.NoError(t, err)

	var opMap map[string]any

	err = json.Unmarshal([]byte(opJson), &opMap)
	require.NoError(t, err)

	userOp, err := userop.New(opMap)
	require.NoError(t, err)

	k, _ := crypto.SKFromInt(big.NewInt(0))
	userOp, _ = UserOpSeal(userOp, big.NewInt(137), &chain.EcdsaKey{SK: k})
	opMap, _ = userOp.ToMap()

	var pmResp map[string]any
	err = pmRpc.Call(&pmResp, "pm_sponsorUserOperation", opMap, DefaultEntryPoint.String(), map[string]string{"type": "payg"})
	require.NoError(t, err)

	for k, v := range pmResp {
		opMap[k] = v
	}

	var reply string
	err = nodeRpc.Call(&reply, "eth_sendUserOperation", opMap, DefaultEntryPoint.String())
	require.NoError(t, err)
	fmt.Printf("Response: %v", reply)

}

var opJson = `{
    "sender": "0xfD6DD93dCc566f6E8C0A5FFb7322B1302c1d2CC0",
    "nonce": "0x0",
    "initCode": "0x9406cc6185a346906296840746125a0e449764545fbfb9cf0000000000000000000000003bf27b2b37345d08a980e273564473bc3744bb1e0000000000000000000000000000000000000000000000000000000000000001",
    "callData": "0xb61d27f600000000000000000000000058a2993a618afee681de23decbcf535a58a080ba000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000044205c28780000000000000000000000003bf27b2b37345d08a980e273564473bc3744bb1e00000000000000000000000000000000000000000000000000000000000f424000000000000000000000000000000000000000000000000000000000",
    "callGasLimit": "0x30d40",
    "verificationGasLimit": "0x6ddd0",
    "preVerificationGas": "0x186a0",
    "maxFeePerGas": "0xfa4e98242",
    "maxPriorityFeePerGas": "0x0",
    "paymasterAndData": "0x",
    "signature": "0x8ef98492cacee9b734e179ebf6423c3cd1b490c33710627a9b5fca64036def231d5138c0b8075739448d6bddac35b55c6cbb8410febb4c47d06f47220b2ca4251c"
}`

func TestOpSubmitDirect(t *testing.T) {

	sk, err := crypto.SKFromHex(os.Getenv("CHAIN_SK"))
	require.NoError(t, err)
	k := &chain.EcdsaKey{SK: sk}

	ec, err := jsonrpc.NewClient(os.Getenv("CHAIN_URL"))
	require.NoError(t, err)

	ep, err := chain.LoadContract(ec, "IEntryPoint.sol/IEntryPoint", k, DefaultEntryPoint)
	require.NoError(t, err)

	var opMap map[string]any

	err = json.Unmarshal([]byte(opJson), &opMap)
	require.NoError(t, err)

	_, err = userop.New(opMap)
	require.NoError(t, err)

	handleError := func(err error) {
		if cerr, ok := err.(*codec.ErrorObject); ok {
			et := errThunk{cerr: cerr}
			if etData, ok := et.ErrorData().(string); ok {
				if strings.HasPrefix(etData, "0xe0cff05f") {
					result, _ := reverts.NewValidationResult(et)
					println(fmt.Sprintf("stake: %v, sig failed: %v", result.SenderInfo.Stake.Int64(), result.ReturnInfo.SigFailed))
				} else {
					revert, _ := reverts.NewFailedOp(et)
					println(revert.Reason)
				}
			}
		}
	}

	err = chain.TxnDoWait(ep.Txn("handleOps", []map[string]any{opMap}, k.Address()))
	if err != nil {
		handleError(err)
	}
	require.NoError(t, err)
}
