import React, { useState, useContext, useEffect } from 'react';
import { Outlet, useLocation, useNavigate } from 'react-router-dom';
import Web3Context from '../Web3Context.js';
import Navbar from './Navbar.jsx';
import Login from './Login.jsx';
import Sidebar from './Sidebar.jsx';
import '../styles/Root.css';

const Root = () => {
  const { loggedIn } = useContext(Web3Context);
  const loco = useLocation();

  const [viewSidebar, setViewSidebar] = useState(false);
  const [currentPath, setCurrentPath] = useState(false);

  const navigate = useNavigate();

  const toggleSidebar = () => {
    setViewSidebar(!viewSidebar);
  }

  const toggleOnScroll = () => {
    if(window.scrollY > 81) {
      setViewSidebar(false);
    }
  }

  useEffect(() => {
    if(loggedIn === false) {
      navigate('/login', { replace: true });
    }
  }, [loggedIn]);

  useEffect(() => {
    setCurrentPath(location.pathname);
    window.addEventListener('scroll', toggleOnScroll);

    return () => window.removeEventListener('scroll', toggleOnScroll);
  }, [])

  useEffect(() => {
    setCurrentPath(loco.pathname);
  }, [loco]);




  return (
    <div>
      <>
        <Navbar
          toggleSidebar={toggleSidebar}
          setViewSidebar={setViewSidebar}
        />
        <Sidebar
          viewSidebar={viewSidebar}
          currentPath={currentPath}
        />
      </>
      <div id='outletBackground' onClick={() => setViewSidebar(false)}>
        <Outlet />
      </div>
    </div>
  );
}

export default Root;