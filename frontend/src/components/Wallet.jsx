import React, { useState, useEffect, useContext } from 'react';
import Web3Context from '../Web3Context.js';
import Web3 from "web3";
import Modal from 'react-modal';
import QRCode from "react-qr-code";
import '../styles/Wallet.css';
import 'dotenv';


const Wallet = () => {

  const { web3, web3auth, provider, setProvider, loggedIn, setLoggedIn } = useContext(Web3Context);

  const [USDCBalance, setUSDCBalance] = useState();
  const [MATICBalance, setMATICBalance] = useState();
  const [SFLUVBalance, setSFLUVBalance] = useState();

  const [walletAddress, setWalletAddress] = useState();
  const [transactionSubscription, setTransactionSubscription] = useState();

  const [depositing, setDepositing] = useState(false);
  const [awaiting, setAwaiting] = useState(false);

  const getERC20Balance = async (contractAddress) => {
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

    const getTokenBalance = async () => {

      const result = await contract.methods.balanceOf(walletAddress).call();

      const formattedResult = web3.utils.fromWei(result, "mwei"); // 29803630.997051883414242659

      return formattedResult;
    }
    const bal = await getTokenBalance();
    return bal;
  }


  const getMATICBalance = async () => {
    if (!provider) {
      console.error("provider not initialized yet");
      return;
    }

    const balance = web3.utils.fromWei(
      await web3.eth.getBalance(walletAddress), // Balance is in wei
      "ether"
    );
    return balance;
  };

  const formatBalance = (balance) => {
    balance = (balance * 1.0).toFixed(2);

    return balance
  }

  const formatMATICBalance = (balance) => {
    balance = (balance * 1.0).toFixed(5);

    return balance
  }


  useEffect(() => {
    if(walletAddress) {

      const getWalletBalances = async () => {

        let USDC = await getERC20Balance('0x3c499c542cEF5E3811e1192ce70d8cC03d5c3359')
          .catch(() => console.error('Couldn\'t get USDC balance'));
        let SFLUV = await getERC20Balance('0x58a2993A618Afee681DE23dECBCF535A58A080BA')
          .catch(() => console.error('Couldn\'t get SFLUV balance'));
        let MATIC = await getMATICBalance()
          .catch(() => console.error('Couldn\'t get MATIC balance'));

        setUSDCBalance(formatBalance(USDC));
        setSFLUVBalance(formatBalance(SFLUV));
        setMATICBalance(formatMATICBalance(MATIC));

      }

      getWalletBalances();
    }
  }, [walletAddress]);

  useEffect(() => {

    if(web3) {
      const getWalletAddress = async () => {
        const address = (await web3.eth.getAccounts())[0];
        setWalletAddress(address);
      }
      getWalletAddress();
    }

  }, [web3]);

  const toggleDepositModal = () => {
    if(depositing) {
      setAwaiting(true);
    }
    setDepositing(!depositing);
  }

  Modal.setAppElement('#app');

  return (
    <>
      <div id='profilePage' className='pageTemplate'>
        <div id='profileSections'>
          <div className='profileSection'>
            <span className='col1'>SFLUV: </span>
            <span className='col2'>{SFLUVBalance}</span>
          </div>
          <div className='profileSection'>
            <span className='col1'>USDC: </span>
            <span className='col2'>{USDCBalance}</span>
          </div>
          <div  className='profileSection capitalized'>
            <span className='col1'>MATIC: </span>
            <span className='col2'>{MATICBalance}</span>
          </div>
        </div>
        {walletAddress && (
          <button
            className='button1 walletButton'
            onClick={toggleDepositModal}
          >
            Deposit
          </button>
        )}
        {awaiting && <p id='warningMessage'>Page refresh required in order for balance changes to be reflected</p>}
        <Modal
          id='depositModal'
          isOpen={depositing}
          onRequestClose={toggleDepositModal}
          className='pageTemplate'
          contentElement={
            (props, contentElement) => (
              <div {...props}>
                <h1>Deposit</h1>
                <QRCode
                  size={256}
                  id='depositQR'
                  value={walletAddress}
                  style={{
                    maxWidth: 'min(50%, 200px)',
                    height: 'auto'
                  }}
                />
                <p id='address'>{walletAddress}</p>
                <p id='warningMessage'><b>Please</b> be sure all transactions are being sent on <b>{process.env.REACT_APP_CHAIN_DISPLAY_NAME}</b></p>
              </div>
            )
          }
          style={{
            overlay: {
              backgroundColor: 'rgb(10, 0, 0, .35)',
            }
          }}
          preventScroll={true}
        >
        </Modal>
      </div>
    </>
  )
}

export default Wallet;