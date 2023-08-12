package erc4337

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/paulgoleary/local-luv-proto/chain"
	"github.com/stackup-wallet/stackup-bundler/pkg/userop"
	"github.com/umbracle/ethgo"
	"github.com/umbracle/ethgo/abi"
	"math/big"
)

// for Mumbai, Polygon Mainnet, ETH mainnet, ...
var DefaultEntryPoint = "0x5ff137d4b0fdcd49dca30c7cf57e578a026d2789"

var abiExec, _ = abi.NewMethod("function execute(address to, uint256 value, bytes data)")

func makeExecute(toAddr ethgo.Address, value *big.Int, m *abi.Method, args ...interface{}) (enc []byte, err error) {
	if enc, err = m.Encode(args); err == nil {
		return abiExec.Encode([]interface{}{toAddr, value, enc})
	}
	return
}

var abiMC, _ = chain.LoadABI("MockCoin.sol/MockCoin")

func UserOpMint(sender, mintAddr, toAddr ethgo.Address, amt *big.Int) (*userop.UserOperation, error) {
	mintMethod := abiMC.GetMethod("mint")
	if callData, err := makeExecute(mintAddr, big.NewInt(0), mintMethod, toAddr, amt); err != nil {
		return nil, err
	} else {
		opData := map[string]any{
			"sender": sender.String(),
			"nonce":  "0x0",
			// TODO: correct init code ...
			"initCode":             "0xe19e9755942bb0bd0cccce25b1742596b8a8250b3bf2c3e700000000000000000000000078d4f01f56b982a3b03c4e127a5d3afa8ebee6860000000000000000000000008b388a082f370d8ac2e2b3997e9151168bd09ff50000000000000000000000000000000000000000000000000000000000000000",
			"callData":             hexutil.Encode(callData),
			"callGasLimit":         "0x0",
			"verificationGasLimit": "0x0",
			"maxFeePerGas":         "0xa862145e",
			"maxPriorityFeePerGas": "0xa8621440",
			"paymasterAndData":     "0x",
			"preVerificationGas":   "0x0",
			"signature":            "0x00",
		}
		return userop.New(opData)
	}
}
