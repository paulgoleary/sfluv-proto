import web3auth from '../web3auth.js';
import React, { useContext, useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import Web3Context from '../Web3Context.js';
import { WALLET_ADAPTERS, IProvider } from '@web3auth/base';
import '../styles/Login.css';
import CircleLoader from 'react-spinners/ClipLoader';
import Logo from '../../assets/SFLUV LOGO-1.svg';



const Root = () => {

  const { provider, setProvider, loggedIn, setLoggedIn } = useContext(Web3Context);

  const navigate = useNavigate();

  const [loggingIn, setLoggingIn] = useState(false);

  useEffect(() => {
    if(loggedIn === true) {
      navigate('/', { replace: true });
    }
  }, [loggedIn]);

  const login = async () => {
    setLoggingIn(true);
    const web3authProvider = await web3auth.connectTo(
      WALLET_ADAPTERS.OPENLOGIN,
      {
        loginProvider: "google",
      }
    )
      .catch((err) => {
        console.error();
        setLoggingIn(false);
      })
    console.log(web3authProvider);
    setProvider(web3authProvider);
    if (web3auth.connected) {
      setLoggedIn(true);
    }
    setLoggingIn(false);
  };

  return (<>
    {!loggingIn && <>
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
    </>}
    {loggingIn &&
      <div
        className='pageTemplate'
        id='loginPage'
        style={{paddingBottom: '35vh', paddingTop: '35vh'}}
      >
          <CircleLoader color='#eb6c6c' loading={true}/>
      </div>
    }
    </>
  )
}

export default Root;
