package erc4337

import (
	"github.com/umbracle/ethgo"
	"github.com/umbracle/ethgo/abi"
	"math/big"
)

// let createAccountTx = await accountFactoryContract.$data().createAccount({ address: owner }, owner, salt);
// let initCode = $abiUtils.encodePacked(['address', 'bytes'], [ accountFactoryContract.address, createAccountTx.data ]);
// let initCodeGas = await this.client.getGasEstimation(entryPointContract.address, createAccountTx);

var DefaultAccountFactory = ethgo.HexToAddress("0x9406Cc6185a346906296840746125a0E44976454")
var DefaultInitSalt = big.NewInt(1)

var createAccountMethod, _ = abi.NewMethod("function createAccount(address owner,uint256 salt) public returns (address ret)")

func MakeDefaultInitCode(owner ethgo.Address) (ret []byte, err error) {
	return makeInitCode(DefaultAccountFactory, owner, DefaultInitSalt)
}

func makeInitCode(factory, owner ethgo.Address, salt *big.Int) (ret []byte, err error) {
	if ret, err = createAccountMethod.Encode([]interface{}{owner, salt}); err != nil {
		return
	}
	ret = append(factory.Bytes(), ret...) // 'packed encoding' ...
	return
}
