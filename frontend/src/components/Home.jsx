import React, { useState, useContext } from 'react';
import Web3 from "web3";
import Web3Context from '../Web3Context.js';

const Home = () => {

  const { web3auth, provider, setProvider, loggedIn, setLoggedIn } = useContext(Web3Context);
  const [signatureMessage, setSignatureMessage] = useState('');

  const getUserInfo = async () => {
    const user = await web3auth.getUserInfo();
    console.log(user);
  };

  const getBalance = async () => {
    if (!provider) {
      console.error("provider not initialized yet");
      return;
    }
    const web3 = new Web3(provider);

    // Get user's Ethereum public address
    const address = (await web3.eth.getAccounts())[0];

    // Get user's balance in ether
    const balance = web3.utils.fromWei(
      await web3.eth.getBalance(address), // Balance is in wei
      "ether"
    );
    console.log(balance);
  };

  const getAccounts = async () => {
    if (!provider) {
      console.error("provider not initialized yet");
      return;
    }
    const web3 = new Web3(provider);

    // Get user's Ethereum public address
    const address = await web3.eth.getAccounts();
    console.log(address);
  };

  const signMessage = async (originalMessage) => {
    if (!provider) {
      console.error("provider not initialized yet");
      return;
    }
    const web3 = new Web3(provider);

    // Get user's Ethereum public address
    const fromAddress = (await web3.eth.getAccounts())[0];

    // Sign the message
    const signedMessage = await web3.eth.personal.sign(
      originalMessage,
      fromAddress,
      'test'
    );

    console.log(signedMessage);
  };

  const getERC20Balance = async (contractAddress) => {
    const web3 = new Web3(provider);

    const balanceOfABI = [
      {
        "constant": true,
        "inputs": [
            {
                "name": "_owner",
                "type": "address"
            }
        ],
        "name": "balanceOf",
        "outputs": [
            {
                "name": "balance",
                "type": "uint256"
            }
        ],
        "payable": false,
        "stateMutability": "view",
        "type": "function"
      },
    ];

    const contract = new web3.eth.Contract(balanceOfABI, contractAddress);
    const address = (await web3.eth.getAccounts())[0];

    const getTokenBalance = async () => {

      const result = await contract.methods.balanceOf(address).call();

      const formattedResult = web3.utils.fromWei(result, "mwei"); // 29803630.997051883414242659

      return formattedResult;
    }
    console.log(await getTokenBalance());
  }

  const logout = async () => {
    await web3auth.logout();
    setProvider(null);
    setLoggedIn(false);
    console.log("logged out");
  };

  return (
    <>
      <div>Home</div>
      <button onClick={logout}>Log Out</button>
      <h2>{process.env.REACT_APP_CHAIN_ID}</h2>
      <button onClick={getUserInfo}>Get User Info</button>
      <button onClick={getAccounts}>Get Accounts</button>
      <button onClick={getBalance}>Get User Balance</button>
      <button onClick={() => getERC20Balance('0x3c499c542cEF5E3811e1192ce70d8cC03d5c3359')}>Get USDC Balance</button>
      <form onSubmit={(e) => {
        e.preventDefault();
        signMessage(signatureMessage);
      }}>
        <input
          onChange={(e) => {
            setSignatureMessage(e.target.value)
          }}
          value={signatureMessage}
        ></input>
        <button type='submit'>Sign Message</button>
      </form>
    </>
  )
}

export default Home;