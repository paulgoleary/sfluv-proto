import React, { useState, useContext } from 'react';
import Web3 from "web3";
import Web3Context from '../Web3Context.js';

const Home = () => {

  const { web3auth, provider, setProvider, loggedIn, setLoggedIn } = useContext(Web3Context);
  const [signatureMessage, setSignatureMessage] = useState('');

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
      <div>Home</div>
      <button onClick={logout}>Log Out</button>
      <h2>{process.env.REACT_APP_CHAIN_ID}</h2>
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
    </>
  )
}

export default Home;