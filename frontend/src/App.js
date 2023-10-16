import React, { useEffect, useState } from 'react';
import Web3  from 'web3';
import auth from './web3auth';
import { UserOperationBuilder } from 'userop';
import links from './links';
import logo from './imgs/SFLUV.svg'

import './App.css';

function App() {
  const [web3, setWeb3] = useState();
  const web3auth = auth.web3auth;
  const l = links();
  const entryPoint = '0x5FF137D4b0FDCD49DcA30c7CF57E578a026d2789';



  const [account, setAccount] = useState('');
  var awaiting = false;

  // const [signatureMessage, setSignatureMessage] = useState('');
  // const [recoverMessage, setRecoverMessage] = useState('');
  // const [recoverSigned, setRecoverSigned] = useState('');
  const [unwrapAddress, setUnwrapAddress] = useState('');
  const [unwrapAmount, setUnwrapAmount] = useState('');
  const [advancedUnwrap, setAdvancedUnwrap] = useState(false);




  const connect = async () => {
    awaiting = true;
    if(!web3auth.connected) {
      await web3auth.connect();
    };
    awaiting = false;
    if (web3auth.provider) {
      const w = new Web3(web3auth.provider)
      setWeb3(w);
    }
  }



  const disconnect = async () => {
    await web3auth.logout();
    setWeb3(undefined);
  }


  const web3authSetup = async () => {
    if(!web3auth.connected) {
      await web3auth.initModal();
    };
    if (!web3auth.connected) {
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
    if (!web3auth.connected && !awaiting) {
      web3authSetup();
    }
    return () => {
      web3auth.loginModal.closeModal();
    }
  }, []);



  const signMessage = async (message: string) => {
    if (web3) {
      const privateKey = await getPrivateKey();
      const signedMessage = web3.eth.accounts.sign(message, '0x' + privateKey);
      return(signedMessage);
    }
  }

  // const recoverAddress = (message: string, signedMessage: string) => {
  //   if (web3) {
  //     const recoveredAddress = web3.eth.accounts.recover(message, signedMessage);
  //   }
  // }

  const fetchData = async (url, postData) => {
    var returnData = undefined;
    await fetch(url, postData)
    .then((res) => {
      if (!res.ok) {
        throw new Error('Failed Fetch');
      }
      const response = res.json();
      if(response){
        return(response);
      } else {
        return(res);
      }
    })
    .then(data => {
      returnData = data;
    })
    return(returnData);
  }

  const signUserOperation = async (ctx) => {
    const hash = ctx.getUserOpHash();
    const signedMessage = await signMessage(hash);
    ctx.op.signature = signedMessage.signature;
  }

  const makeUserOp = async (amount, address) => {
    // const firstFetchData = await fetchData(l.initial + 'owner=' + account);
    const link = l.unwrap +
    'target=58a2993A618Afee681DE23dECBCF535A58A080BA&to=' +
    (address ? address : account) +
    '&amount=' +
    amount +
    '&owner=' +
    account;
    const userOpData = await fetchData(link);
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
    return(op);

  }

  const unwrap = async (amount, address) => {
    amount = amount * 1000000;
    const op = await makeUserOp(amount, address);
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
      {!web3auth.connected && <img src={logo} alt='SFLUV' className='logo' id='logged-out-logo'></img>}
      {web3auth.connected && <div id='unwrap-tool-title'>
        <img src={logo} alt='SFLUV' className='logo' id='logged-in-logo'></img>
        <div id='unwrap-tool-title-text'><h1 className='title'>Unwrap Tool</h1></div>
      </div>}
      {/* <p id='address'>{account}</p> */}
      {web3auth.connected && <div>
        <form
          onSubmit={(e) => {
            e.preventDefault();
            unwrap(unwrapAmount, unwrapAddress);
          }}
          id='unwrap-form'>
          <input
            type='number'
            step='.000001'
            onWheel={e => {
              e.currentTarget.blur();
            }}
            id='unwrap-amount-input'
            className='single-input'
            onChange={(e) => {
              setUnwrapAmount(e.target.value);
            }}
            placeholder='AMOUNT'
          ></input>
          <br/>
          <button id='unwrap-button' className='submit-button'>Unwrap</button>
          <br/>
          <button
            className='dropdown-button'
            id='advanced-unwrap-dropdown'
            onClick={(e) => {
              e.preventDefault();
              if (advancedUnwrap) {
                setAdvancedUnwrap(false);
              } else {
                setAdvancedUnwrap(true);
              }
            }}
          >
            {advancedUnwrap ? 'v Advanced' : '> Advanced'}
          </button>
          <br/>
          {advancedUnwrap && <div>
            <input
            type='text'
            id='unwrap-address-input'
            className='single-input'
            onChange={(e) => {
              setUnwrapAddress(e.target.value);
            }}
            placeholder='ADDRESS'
            ></input>
            <br/>
          </div>}
        </form>
        <button onClick={disconnect}>Disconnect</button>
        {/* <button onClick={getAccounts}>Get Accounts</button>
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
        </form> */}
      </div>}
      {!web3auth.connected && <div>
        <button onClick={connect}>Connect</button>
      </div>}
      {/* <form onSubmit={(e) => {
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
        </form> */}
    </div>
  );
}

export default App;
