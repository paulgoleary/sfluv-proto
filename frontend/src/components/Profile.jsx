import React, { useState, useContext, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import Web3Context from '../Web3Context.js';

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

  return (
    <>
      <div>
        <div>{userName}</div>
        <div>{userEmail}</div>
        <div style={{textTransform: 'capitalize'}}>{loggedInWith}</div>
        <button onClick={() => {console.log('working')}} />
      </div>
    </>
  )
}

export default Profile;