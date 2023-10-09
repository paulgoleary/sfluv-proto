import React, { useEffect, useState } from 'react';
import Web3  from 'web3';
import auth from './web3auth';
import { UserOperationBuilder } from 'userop';
import links from './links';

import './App.css';

function App() {
  const [web3, setWeb3] = useState();
  const web3auth = auth.web3auth;
  const l = links();
  const entryPoint = '0x5FF137D4b0FDCD49DcA30c7CF57E578a026d2789';


  const [connected, setConnected] = useState(false);
  const [account, setAccount] = useState('');

  const [signatureMessage, setSignatureMessage] = useState('');
  const [recoverMessage, setRecoverMessage] = useState('');
  const [recoverSigned, setRecoverSigned] = useState('');




  const connect = async () => {
    await web3auth.connect();
    setConnected(web3auth.connected);
    if (web3auth.provider) {
      const w = new Web3(web3auth.provider)
      setWeb3(w);

      console.log(w);
    }
  }



  const disconnect = async () => {
    await web3auth.logout();
    setWeb3(undefined);
    setConnected(await web3auth.connected);
  }


  const web3authSetup = async () => {
    await web3auth.initModal();
    setConnected(web3auth.connected);
    if (!connected) {
      connect();
    };
  }

  const getPrivateKey = async () => {
    const privateKey = await web3auth.provider.request({
      method: "eth_private_key"
    });
    return(privateKey);
  }

  const getAccounts = async () => {
    if (web3) {
      const accounts = (await web3.eth.getAccounts());
      return accounts;
    }
  }

  const accountSetter = async () => {
    if (web3) {
      const a = await getAccounts();
      setAccount(a[0]);
    }
  }

  useEffect(() => {
    accountSetter();
  }, [web3]);


  useEffect(() => {
    if (!connected) {
      web3authSetup();
    }
  }, []);



  const signMessage = async (message: string) => {
    if (web3) {
      const privateKey = await getPrivateKey();
      const signedMessage = web3.eth.accounts.sign(message, '0x' + privateKey);
      // console.log(signedMessage);
      return(signedMessage);
    }
  }

  const recoverAddress = (message: string, signedMessage: string) => {
    if (web3) {
      const recoveredAddress = web3.eth.accounts.recover(message, signedMessage);
      console.log(recoveredAddress);
    }
  }

  const fetchData = async (url, postData) => {
    var returnData = undefined;
    await fetch(url, postData)
    .then((res) => {
      if (!res.ok) {
        throw new Error('Failed Fetch');
      }
      const response = res.json();
      // console.log(response);
      // console.log(res);
      if(response){
        return(response);
      } else {
        return(res);
      }
    })
    .then(data => {
      console.log(data);
      returnData = data;
    })
    return(returnData);
  }

  const signUserOperation = async (ctx) => {
    const hash = ctx.getUserOpHash();
    // console.log(ctx);
    console.log('Hash:' + hash);
    const signedMessage = await signMessage(hash);
    // console.log(signedMessage);
    ctx.op.signature = signedMessage.signature;
  }

  const makeUserOp = async () => {
    const firstFetchData = await fetchData(l.initial + 'owner=' + account);
    const userOpData = await fetchData(
      l.unwrap +
      'target=2791Bca1f2de4661ED88A30C99A7a9449Aa84174&to=' +
      account +
      '&amount=100000000&owner=' +
      account
    );
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
      .useMiddleware(signUserOperation)
    ;
    let op = await builder.buildOp(entryPoint , 137);
    console.log(op);
    return(op);

  }

  const unwrap = async () => {
    const op = await makeUserOp();
    console.log(op);
    const postData = { entryPoint, op };
    console.log(postData);
    const post = await fetchData(l.post, {
      method: 'POST',
      headers: {'Content-Type': 'application/json', 'Accept': 'application/json' },
      body: JSON.stringify(postData)
    });
  }

  return (
    <div className="App">
      <p>Hi</p>
      {connected && <div>

        <button onClick={disconnect}>Disconnect</button>
        <br/>
        <button onClick={getAccounts}>Get Accounts</button>
        <form onSubmit={(e) => {
          e.preventDefault();
          signMessage(signatureMessage);
        }}>
          <input
            type='text'
            value={signatureMessage}
            onChange={(e) => setSignatureMessage(e.target.value)}>
          </input>
          <button>Sign Message</button>
        </form>
        <button onClick={unwrap}>Unwrap</button>

      </div>}
      {!connected && <div>

        <button onClick={connect}>Connect</button>

      </div>}
      <form onSubmit={(e) => {
          e.preventDefault();
          recoverAddress(recoverMessage, recoverSigned);
        }}>
          <p>Message:</p>
          <input
            type='text'
            value={recoverMessage}
            onChange={(e) => setRecoverMessage(e.target.value)}>
          </input>
          <p>Signed:</p>
          <input
            type='text'
            value={recoverSigned}
            onChange={(e) => setRecoverSigned(e.target.value)}>
          </input>
          <br></br>
          <button>Recover Address</button>
        </form>
        <p>{account}</p>
    </div>
  );
}

export default App;
