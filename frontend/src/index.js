import React, { useState, useEffect, createContext } from 'react';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import Root from './components/Root.jsx';
import ErrorPage from './components/ErrorPage.jsx';
import Login from './components/Login.jsx';
import Home from './components/Home.jsx';
import web3auth from './web3auth.js';
import Web3Context from './Web3Context.js';
import { createRoot } from 'react-dom/client';
import './style.css';
import 'dotenv';


const App = () => {

  const [provider, setProvider] = useState();
  const [loggedIn, setLoggedIn] = useState();

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

    init();
  }, []);

  const router = createBrowserRouter([
    {path: '/', element: <Root />, errorElement: <ErrorPage />, children: [
      {path: '/', element: <Home />}
    ]}
  ])

  return (
    <Web3Context.Provider value={{ provider, setProvider, loggedIn, setLoggedIn, web3auth }}>
      <RouterProvider router={router} />
    </Web3Context.Provider>
  );
}

createRoot(document.getElementById('app')).render(<App />);
