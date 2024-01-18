import React, { useState, useContext, useEffect } from 'react';
import { Outlet } from 'react-router-dom';
import Web3Context from '../Web3Context.js';
import Navbar from './Navbar.jsx';
import Login from './Login.jsx';
import Sidebar from './Sidebar.jsx';
import '../styles/Root.css';

const Root = () => {
  const { loggedIn } = useContext(Web3Context);

  const [viewSidebar, setViewSidebar] = useState(false);
  const [currentPath, setCurrentPath] = useState(false);

  const toggleSidebar = () => {
    setViewSidebar(!viewSidebar);
  }

  useEffect(() => {
    setCurrentPath(location.pathname)
  }, [])


  return (
    <div>
      {loggedIn === false &&
        <Login />
      }
      {loggedIn &&
      <>
        <Navbar
          toggleSidebar={toggleSidebar}
          setCurrentPath={setCurrentPath}
          setViewSidebar={setViewSidebar}
        />
        <Sidebar
          viewSidebar={viewSidebar}
          currentPath={currentPath}
          setCurrentPath={setCurrentPath}
        />
        <div id='outletBackground' onClick={() => setViewSidebar(false)}>
          <Outlet />
        </div>
      </>
      }
    </div>
  );
}

export default Root;