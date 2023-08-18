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

// 0x80FD6A0a454045d3E24B5BAa19b9066ae01a0b09
// 0x39C716D8c6E4D3B45Bc9e60f5C12378433668588
var MockMumbaiAddr = ethgo.HexToAddress("0x80FD6A0a454045d3E24B5BAa19b9066ae01a0b09")
var LuvMumbaiAddr = ethgo.HexToAddress("0x39C716D8c6E4D3B45Bc9e60f5C12378433668588")

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
