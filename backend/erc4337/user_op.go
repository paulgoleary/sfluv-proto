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
var DefaultEntryPoint = ethgo.HexToAddress("0x5ff137d4b0fdcd49dca30c7cf57e578a026d2789")

var abiExec, _ = abi.NewMethod("function execute(address to, uint256 value, bytes data)")

func makeExecute(toAddr ethgo.Address, value *big.Int, m *abi.Method, args ...interface{}) (enc []byte, err error) {
	if enc, err = m.Encode(args); err == nil {
		return abiExec.Encode([]interface{}{toAddr, value, enc})
	}
	return
}

var abiMC, _ = chain.LoadABI("MockCoin.sol/MockCoin")

var DefaultInitCodeGas = big.NewInt(300_000)
var DefaultMintGasLimit = big.NewInt(200_000)

func UserOpMint(owner, sender, mintAddr, toAddr ethgo.Address, amt *big.Int) (*userop.UserOperation, error) {
	mintMethod := abiMC.GetMethod("mint")
	if callData, err := makeExecute(mintAddr, big.NewInt(0), mintMethod, toAddr, amt); err != nil {
		return nil, err
	} else {

		initCode, err := MakeInitCode(DefaultAccountFactory, owner, big.NewInt(1))
		if err != nil {
			return nil, err
		}

		vGasLimit := new(big.Int).Add(big.NewInt(150_000), DefaultInitCodeGas)
		opData := map[string]any{
			"sender":               sender.String(),
			"nonce":                "0x0",
			"initCode":             hexutil.Encode(initCode),
			"callData":             hexutil.Encode(callData),
			"callGasLimit":         DefaultMintGasLimit,
			"verificationGasLimit": vGasLimit,
			"maxFeePerGas":         big.NewInt(2_000_000_000),
			"maxPriorityFeePerGas": "0x0",
			"paymasterAndData":     "0x",
			"preVerificationGas":   big.NewInt(100_000),
			"signature":            "0x00",
		}
		return userop.New(opData)
	}
}
