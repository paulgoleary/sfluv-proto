import { UserOperationBuilder } from 'userop';
import { useContext } from 'react';
import Web3Context from '../Web3Context.js';
import 'dotenv';

const signUserOp = async (web3, privateKey, userOpData, account) => {

  const signMiddleware = async (ctx) => {

    let hash = ctx.getUserOpHash();

    const signedMessage = await web3.eth.accounts.sign(hash, '0x' + privateKey);

    const recovery = await web3.eth.accounts.recover(hash, signedMessage.signature)

    ctx.op.signature = signedMessage.signature;
  }

  const builder = new UserOperationBuilder()
    .setNonce(userOpData.nonce)
    .setSender(userOpData.sender)
    .setCallData(userOpData.callData)
    .setCallGasLimit(userOpData.callGasLimit)
    .setInitCode(userOpData.initCode)
    .setMaxFeePerGas(userOpData.maxFeePerGas)
    .setMaxPriorityFeePerGas(userOpData.maxPriorityFeePerGas)
    .setPaymasterAndData(userOpData.paymasterAndData)
    .setPreVerificationGas(userOpData.preVerificationGas)
    .setVerificationGasLimit(userOpData.verificationGasLimit)
    .useMiddleware(signMiddleware);

  const entryPoint = process.env.REACT_APP_ENTRYPOINT;
  const chainId = process.env.REACT_APP_CHAIN_ID;

  const signedOp = await builder.buildOp(entryPoint, chainId)

  return signedOp;
}

export default signUserOp;

