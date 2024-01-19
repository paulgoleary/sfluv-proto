import React from 'react';
import { Sidebar, Menu, MenuItem, SubMenu } from 'react-pro-sidebar';
import { useNavigate } from 'react-router-dom';
import '../styles/Sidebar.css';

const SidebarComponent = ({ viewSidebar, currentPath }) => {

  const navigate = useNavigate();


  return (
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
      onBackdropClick={() => console.log('test')}
    >
      <Menu>
        <MenuItem
          className={'sideBarItem'}
          active={true}
          onClick={() => navigate('/wallet')}
          style={currentPath === '/wallet' ? {
            backgroundColor: '#eb6c6c'
          } : {}}
        >
          Wallet
        </MenuItem>
        <MenuItem
          className={'sideBarItem'}
          onClick={() => navigate('/unwrap')}
          style={currentPath === '/unwrap' ? {
            backgroundColor: '#eb6c6c'
          } : {}}
        >
          Unwrap Tool
        </MenuItem>
        <MenuItem
          className={'sideBarItem'}
          onClick={() => navigate('/profile')}
          style={currentPath === '/profile' ? {
            backgroundColor: '#eb6c6c'
          } : {}}
        >
          My Profile
        </MenuItem>
      </Menu>
    </Sidebar>
  )
}

export default SidebarComponent;