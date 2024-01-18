import React, { useState, useContext } from 'react';
import { Outlet, Link } from 'react-router-dom';
import { Sidebar, Menu, MenuItem, SubMenu } from 'react-pro-sidebar';
import Web3Context from '../Web3Context.js';
import Navbar from './Navbar.jsx';
import Login from './Login.jsx';
import '../styles/Sidebar.css';

const Root = () => {
  const { loggedIn } = useContext(Web3Context);

  const [viewSidebar, setViewSidebar] = useState(false);

  const toggleSidebar = () => {
    setViewSidebar(!viewSidebar);
  }




  return (
    <>
      {loggedIn === false &&
        <Login />
      }
      {loggedIn &&
      <>
        <Navbar toggleSidebar={toggleSidebar}/>
        <Sidebar
          collapsed={!viewSidebar}
          collapsedWidth={'0px'}
          width={'min(40%, 250px)'}
          rtl={true}
          id='sideBar'
          transitionDuration={250}
          style={{
            left: viewSidebar ? 'calc(100% - min(40%, 250px))' : '100%'
          }}
          backgroundColor={'white'}
        >
          <Menu>
            <MenuItem
              className={'sideBarItem'}
              component={<Link to='/wallet'/>}
            >
              Wallet
            </MenuItem>
            <MenuItem
              className={'sideBarItem'}
              component={<Link to='/unwrap'/>}
            >
              Unwrap Tool
            </MenuItem>
            <MenuItem
              className={'sideBarItem'}
              component={<Link to='/profile'/>}
            >
             My Profile
            </MenuItem>
          </Menu>
        </Sidebar>
        <Outlet />
      </>
      }
    </>
  );
}

export default Root;