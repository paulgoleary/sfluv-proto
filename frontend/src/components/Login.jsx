import web3auth from '../web3auth.js';
import React, { useContext, useState } from 'react';
import Web3Context from '../Web3Context.js';
import { WALLET_ADAPTERS, IProvider } from '@web3auth/base';
import '../styles/Login.css';
import Logo from '../../assets/SFLUV LOGO-1.svg';



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
    <div
      className='pageTemplate'
      id='loginPage'
    >
      <h1>Community Portal</h1>
      <img
        src={Logo}
        alt={'SFLUV'}
        id='loginLogo'
      />
      <button
        className='button2'
        id='loginButton'
        onClick={login}
      >
        Log In with Google
      </button>
    </div>
  )
}

export default Root;
