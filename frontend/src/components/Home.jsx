import React, { useState, useContext, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import Web3 from "web3";
import Web3Context from '../Web3Context.js';
import '../styles/Home.css'

const Home = () => {

  const { web3auth, provider, setProvider, loggedIn, setLoggedIn } = useContext(Web3Context);
  const [signatureMessage, setSignatureMessage] = useState('');
  const [firstName, setFirstName] = useState('');

  const navigate = useNavigate();

  useEffect(() => {
    if(loggedIn) {
      const getUserInfo = async () => {
        try{
          const user = await web3auth.getUserInfo();
          setFirstName(user.name.split(' ')[0]);
        } catch {
          navigate('/error');
        }
      };
      getUserInfo();
    }
  }, [loggedIn]);

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



  return (
    <>
      <div className='pageTemplate' id='homepageMessageBox'>
        <h1>Welcome {firstName}!</h1>
        <div id='homepageMessage'>The SFLUV Community Portal is your one stop shop for community voting, unwrapping your LUV, and any technical support you may need!</div>
        <button
          className='homepageButton button1'
          onClick={() => navigate('/wallet')}
        >
          My Wallet
        </button>
        <button
          className='homepageButton button2'
          onClick={() => navigate('/unwrap')}
        >
          Unwrap
        </button>
        <br />
        <br />
      </div>
    </>
  )
}

export default Home;