import React from 'react';
import { Sidebar, Menu, MenuItem, SubMenu } from 'react-pro-sidebar';
import { useNavigate } from 'react-router-dom';
import '../styles/Sidebar.css';

const SidebarComponent = ({ viewSidebar, currentPath, setCurrentPath }) => {

  const navigate = useNavigate();

  const navigateAndUpdatePath = (path) => {
    setCurrentPath(path);
    navigate(path);
  }

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
          onClick={() => navigateAndUpdatePath('/wallet')}
          style={{
            backgroundColor: currentPath === '/wallet' ? '#eb6c6c' : 'whitesmoke'
          }}
        >
          Wallet
        </MenuItem>
        <MenuItem
          className={'sideBarItem'}
          onClick={() => navigateAndUpdatePath('/unwrap')}
          style={{
            backgroundColor: currentPath === '/unwrap' ? '#eb6c6c' : 'whitesmoke'
          }}
        >
          Unwrap Tool
        </MenuItem>
        <MenuItem
          className={'sideBarItem'}
          onClick={() => navigateAndUpdatePath('/profile')}
          style={{
            backgroundColor: currentPath === '/profile' ? '#eb6c6c' : 'whitesmoke'
          }}
        >
          My Profile
        </MenuItem>
      </Menu>
    </Sidebar>
  )
}

export default SidebarComponent;