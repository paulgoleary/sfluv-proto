import React, { useState, useContext, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import Web3Context from '../Web3Context.js';
import '../styles/Profile.css';

const Profile = () => {
  const { web3auth, provider, setProvider, loggedIn, setLoggedIn } = useContext(Web3Context);

  const navigate = useNavigate();

  const [user, setUser] = useState();
  const [userName, setUserName] = useState();
  const [userEmail, setUserEmail] = useState();
  const [loggedInWith, setLoggedInWith] = useState();



  useEffect(() => {
    const getUserInfo = async () => {
      try{
        const user = await web3auth.getUserInfo();
        setUser(true);
        setUserName(user.name);
        setUserEmail(user.email);
        setLoggedInWith(user.typeOfLogin);
      } catch {
        navigate('/error');
      }
    };
    getUserInfo();
  }, []);


  const logout = async () => {
    await web3auth.logout();
    setProvider(null);
    setLoggedIn(false);
    navigate('/');
    console.log("logged out");
  };


  return (
    <div id='profilePage' className='pageTemplate'>
      <div id='profileSections'>
        <div className='profileSection'>
          <span className='col1'>Name: </span>
          <span className='col2'>{userName}</span>
        </div>
        <div className='profileSection'>
          <span className='col1'>Email: </span>
          <span className='col2'>{userEmail}</span>
        </div>
        <div  className='profileSection capitalized'>
          <span className='col1'>Login Provider: </span>
          <span className='col2'>{loggedInWith}</span>
        </div>
      </div>
      <button id='logoutButton' className='button2' onClick={logout}>Log Out</button>
    </div>
  )
}

export default Profile;