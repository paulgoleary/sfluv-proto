import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import Logo from '../../assets/LUV_LOGO.svg'
import { MdDehaze } from "react-icons/md";
import '../styles/Navbar.css';

const Navbar = ({ toggleSidebar }) => {

  const navigate = useNavigate();

  return (
    <div id='navBar'>
        <img id='navbarLogo' src={Logo} onClick={() => navigate('/')}/>
        <button onClick={toggleSidebar} id='navMenuButton'>
          <MdDehaze id='navMenuIcon'/>
        </button>

    </div>
  );
}

export default Navbar;