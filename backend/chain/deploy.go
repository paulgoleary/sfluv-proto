package chain

import (
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"github.com/paulgoleary/local-luv-proto/crypto"
	"github.com/umbracle/ethgo"
	"github.com/umbracle/ethgo/abi"
	"github.com/umbracle/ethgo/compiler"
	"github.com/umbracle/ethgo/contract"
	"github.com/umbracle/ethgo/jsonrpc"
	"os"
	"path/filepath"
	"strings"
)

var MockMumbaiAddr = ethgo.HexToAddress("0x834F9b26Cc7C806c6F9f31697C4B1C20A1bB83b6")
var LuvMumbaiAddr = ethgo.HexToAddress("0xdcF0C250a68B835cb0379381F28F45732746F177")

type jsonBytecode struct {
	Object string `json:"object"`
}
type jsonArtifact struct {
	Bytecode jsonBytecode    `json:"bytecode"`
	Abi      json.RawMessage `json:"abi"`
}

func getBuildArtifact(name string) (art *compiler.Artifact, err error) {

	if !strings.HasSuffix(name, ".json") {
		name = name + ".json"
	}
	var jsonBytes []byte
	if jsonBytes, err = os.ReadFile(filepath.Join("../../contracts/out", name)); err != nil {
		return
	}
	var jart jsonArtifact
	if err = json.Unmarshal(jsonBytes, &jart); err != nil {
		return
	}

	bc := jart.Bytecode.Object
	if !strings.HasPrefix(bc, "0x") {
		bc = "0x" + bc
	}
	art = &compiler.Artifact{
		Abi: string(jart.Abi),
		Bin: bc,
	}
	return
}

type EcdsaKey struct {
	SK *ecdsa.PrivateKey
}

func (e *EcdsaKey) Address() ethgo.Address {
	return ethgo.Address(crypto.PubKeyToAddress(&e.SK.PublicKey))
}

func (e *EcdsaKey) Sign(hash []byte) ([]byte, error) {
	return crypto.Sign(e.SK, hash)
}

var _ ethgo.Key = &EcdsaKey{}

func LoadContract(ec *jsonrpc.Client, name string, withKey ethgo.Key, addr ethgo.Address) (loaded *contract.Contract, err error) {
	var art *compiler.Artifact
	if art, err = getBuildArtifact(name); err != nil {
		return
	}
	var theAbi *abi.ABI
	if theAbi, err = abi.NewABI(art.Abi); err != nil {
		return
	}

	loaded = contract.NewContract(addr, theAbi,
		contract.WithJsonRPC(ec.Eth()),
		contract.WithSender(withKey),
	)

	return
}

func LoadABI(name string) (loaded *abi.ABI, err error) {
	var art *compiler.Artifact
	if art, err = getBuildArtifact(name); err != nil {
		return
	}
	return abi.NewABI(art.Abi)
}

func deployContract(ec *jsonrpc.Client, name string, withKey ethgo.Key, args []interface{}) (deployed *contract.Contract, addr ethgo.Address, err error) {
	var art *compiler.Artifact
	if art, err = getBuildArtifact(name); err != nil {
		return
	}
	var theAbi *abi.ABI
	if theAbi, err = abi.NewABI(art.Abi); err != nil {
		return
	}

	var rcpt *ethgo.Receipt
	var artBin []byte
	if artBin, err = hex.DecodeString(strings.TrimPrefix(art.Bin, "0x")); err != nil {
		return
	}
	var txn contract.Txn
	if txn, err = contract.DeployContract(theAbi, artBin, args,
		contract.WithJsonRPC(ec.Eth()), contract.WithSender(withKey)); err != nil {
		return
	} else {
		if err = txn.Do(); err != nil {
			return
		}
		if rcpt, err = txn.Wait(); err != nil {
			return
		}
	}

	deployed = contract.NewContract(rcpt.ContractAddress, theAbi,
		contract.WithJsonRPC(ec.Eth()),
		contract.WithSender(withKey),
	)
	addr = rcpt.ContractAddress

	return
}
