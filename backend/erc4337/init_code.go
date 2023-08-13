package erc4337

import (
	"github.com/paulgoleary/local-luv-proto/chain"
	"github.com/umbracle/ethgo"
	"math/big"
)

// let createAccountTx = await accountFactoryContract.$data().createAccount({ address: owner }, owner, salt);
// let initCode = $abiUtils.encodePacked(['address', 'bytes'], [ accountFactoryContract.address, createAccountTx.data ]);
// let initCodeGas = await this.client.getGasEstimation(entryPointContract.address, createAccountTx);

var DefaultAccountFactory = ethgo.HexToAddress("0x9406Cc6185a346906296840746125a0E44976454")

var abiAccountFactory, _ = chain.LoadABI("SimpleAccountFactory.sol/SimpleAccountFactory")

func MakeInitCode(factory, owner ethgo.Address, salt *big.Int) (ret []byte, err error) {
	createMethod := abiAccountFactory.GetMethod("createAccount")
	if ret, err = createMethod.Encode([]interface{}{owner, salt}); err != nil {
		return
	}
	ret = append(factory.Bytes(), ret...) // 'packed encoding' ...
	return
}
