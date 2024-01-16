import React, { useState, useContext } from 'react';
import { Outlet } from 'react-router-dom';
import Web3Context from '../Web3Context.js';
import Navbar from './Navbar.jsx';
import Login from './Login.jsx';

const Root = () => {
  const { loggedIn } = useContext(Web3Context);

  return (
    <>
      {!loggedIn &&
        <Login />
      }
      {loggedIn &&
      <>
        <Navbar />
        <Outlet />
      </>
      }
    </>
  );
}

export default Root;