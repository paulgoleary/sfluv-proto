import React, { useState, useEffect } from 'react';
import web3auth from './web3auth.js';
import { createRoot } from 'react-dom/client';
import { WALLET_ADAPTERS } from '@web3auth/base';
import Web3 from "web3";
import './style.css';
import 'dotenv';


const App = () => {

  const [provider, setProvider] = useState();
  const [loggedIn, setLoggedIn] = useState();
  const [signatureMessage, setSignatureMessage] = useState('');

  useEffect(() => {
    const init = async () => {
      try {
        await web3auth.init();
        setProvider(web3auth.provider);

        if (web3auth.connected) {
          setLoggedIn(true);
        }
      } catch (error) {
        console.error(error);
      }
    }

    init()
    .then(() => console.log(web3auth));
  }, []);

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

  const getUserInfo = async () => {
    const user = await web3auth.getUserInfo();
    console.log(user);
  };

  const getBalance = async () => {
    if (!provider) {
      console.error("provider not initialized yet");
      return;
    }
    const web3 = new Web3(provider);

    // Get user's Ethereum public address
    const address = (await web3.eth.getAccounts())[0];

    // Get user's balance in ether
    const balance = web3.utils.fromWei(
      await web3.eth.getBalance(address), // Balance is in wei
      "ether"
    );
    console.log(balance);
  };

  const getAccounts = async () => {
    if (!provider) {
      console.error("provider not initialized yet");
      return;
    }
    const web3 = new Web3(provider);

    // Get user's Ethereum public address
    const address = await web3.eth.getAccounts();
    console.log(address);
  };

  const signMessage = async (originalMessage) => {
    if (!provider) {
      console.error("provider not initialized yet");
      return;
    }
    const web3 = new Web3(provider);

    // Get user's Ethereum public address
    const fromAddress = (await web3.eth.getAccounts())[0];

    // Sign the message
    const signedMessage = await web3.eth.personal.sign(
      originalMessage,
      fromAddress,
      'test'
    );

    console.log(signedMessage);
  };

  const logout = async () => {
    await web3auth.logout();
    setProvider(null);
    setLoggedIn(false);
    console.log("logged out");
  };

  return (
    <>
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
          <button onClick={logout}>Log Out</button>
        </>}
      </div>
    </>
  );
}

createRoot(document.getElementById('app')).render(<App />);
