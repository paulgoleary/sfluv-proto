import React, { useState, useEffect, createContext } from 'react';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import Root from './components/Root.jsx';
import ErrorPage from './components/ErrorPage.jsx';
import Profile from './components/Profile.jsx';
import Wallet from './components/Wallet.jsx';
import Login from './components/Login.jsx';
import UnwrapTool from './components/UnwrapTool.jsx';
import Home from './components/Home.jsx';
import web3auth from './web3auth.js';
import Web3Context from './Web3Context.js';
import Web3 from 'web3';
import { createRoot } from 'react-dom/client';
import './style.css';
import 'dotenv';


const App = () => {

  const [provider, setProvider] = useState();
  const [loggedIn, setLoggedIn] = useState();
  const [web3, setWeb3] = useState();

  useEffect(() => {
    const init = async () => {
      try {
        await web3auth.init();
        setProvider(web3auth.provider);

        if (web3auth.connected) {
          setLoggedIn(true);
        } else {
          setLoggedIn(false);
        }
      } catch (error) {
        console.error(error);
      }
    }

    init();
  }, []);


  useEffect(() => {
    if(provider) {
      setWeb3(new Web3(provider));
    } else {
      setWeb3(undefined);
    }
  }, [provider])

  const router = createBrowserRouter([
    {path: '/', element: <Root />, errorElement: <ErrorPage />, children: [
      {path: '/', element: <Home />},
      {path: '/profile', element: <Profile />},
      {path: '/wallet', element: <Wallet />},
      {path: '/unwrap', element: <UnwrapTool />},
      {path: '/error', element: <ErrorPage />},
    ]}
  ])

  return (
    <Web3Context.Provider value={{ provider, setProvider, loggedIn, setLoggedIn, web3auth, web3 }}>
      <RouterProvider router={router} />
    </Web3Context.Provider>
  );
}

createRoot(document.getElementById('app')).render(<App />);
