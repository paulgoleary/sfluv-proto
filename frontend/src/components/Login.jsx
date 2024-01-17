import web3auth from '../web3auth.js';
import React, { useContext, useState } from 'react';
import Web3Context from '../Web3Context.js';
import { WALLET_ADAPTERS, IProvider } from '@web3auth/base';


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
