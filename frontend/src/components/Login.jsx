import web3auth from '../web3auth.js';
import React, { useContext, useState } from 'react';
import Web3Context from '../Web3Context.js';
import { WALLET_ADAPTERS } from '@web3auth/base';


const Root = () => {

  const { provider, setProvider, loggedIn, setLoggedIn } = useContext(Web3Context);

  const login = async () => {
    const web3authProvider = await web3auth.connectTo(
      WALLET_ADAPTERS.OPENLOGIN,
      {
        loginProvider: "google",
      }
    );
    console.log(web3authProvider);
    setProvider(web3authProvider);
    if (web3auth.connected) {
      setLoggedIn(true);
    }
  };

  return (
    <>
      <button onClick={login}>Login</button>
    </>
  )
}

export default Root;

{/* <>
      <div>
        <h1>Hello World!</h1>
        <h2>{process.env.REACT_APP_CHAIN_ID}</h2>
        {!loggedIn && <button onClick={login}>Login</button>}
        {loggedIn && <>
          <button onClick={getUserInfo}>Get User Info</button>
          <button onClick={getAccounts}>Get Accounts</button>
          <button onClick={getBalance}>Get User Balance</button>
          <form onSubmit={(e) => {
            e.preventDefault();
            signMessage(signatureMessage);
          }}>
            <input
              onChange={(e) => {
                setSignatureMessage(e.target.value)
              }}
              value={signatureMessage}
            ></input>
            <button type='submit'>Sign Message</button>
          </form>
        </>}
      </div>
    </> */}