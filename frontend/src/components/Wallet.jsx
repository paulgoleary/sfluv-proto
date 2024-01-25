import React, { useState, useEffect, useContext } from 'react';
import { AiFillCheckCircle } from "react-icons/ai";
import Web3Context from '../Web3Context.js';
import Web3 from "web3";
import Modal from 'react-modal';
import QRCode from "react-qr-code";
import links from '../../utils/links.js';
import addresses from '../../utils/addresses';
import axios from 'axios';
import getPrivateKey from '../models/getPrivateKey.js';
import signUserOp from '../models/signUserOp.js';
import getERC20Balance from '../models/getERC20Balance';
import { QrReader } from 'react-qr-reader';
import CircleLoader from 'react-spinners/ClipLoader';
import '../styles/Wallet.css';
import 'dotenv';


const Wallet = () => {

  const { web3, web3auth, provider, setProvider, loggedIn, setLoggedIn } = useContext(Web3Context);

  const [USDCBalance, setUSDCBalance] = useState();
  const [USDCeBalance, setUSDCeBalance] = useState();
  const [MATICBalance, setMATICBalance] = useState();
  const [SFLUVBalance, setSFLUVBalance] = useState();

  const [walletAddress, setWalletAddress] = useState();
  const [ownerAccountAddress, setOwnerAccountAddress] = useState();
  const [transactionSubscription, setTransactionSubscription] = useState();

  const [depositing, setDepositing] = useState(false);
  const [awaiting, setAwaiting] = useState(false);

  const [sending, setSending] = useState(false);
  const [sendingStage, setSendingStage] = useState(0);
  const [sendingAddress, setSendingAddress] = useState('');
  const [invalidAddress, setInvalidAddress] = useState(false);
  const [coinToSend, setCoinToSend] = useState();
  const [coinToSendName, setCoinToSendName] = useState();
  const [coinToSendBalance, setCoinToSendBalance] = useState();
  const [sendingUserOp, setSendingUserOp] = useState();
  const [sendingAmount, setSendingAmount] = useState();
  const [sent, setSent] = useState();

  const [loading, setLoading] = useState(false);

  const [tooHigh, setTooHigh] = useState();
  const [disableSubmit, setDisableSubmit] = useState();

  const coins = [
    {name: 'SFLUV', address: addresses.sfluv, balance: SFLUVBalance},
    {name: 'USDC', address: addresses.usdc, balance: USDCBalance},
    {name: 'USDC.e', address: addresses.usdce, balance: USDCeBalance},
    // {name: 'MATIC', address: addresses.matic, balance: MATICBalance}
  ]

  useEffect(() => {
    if(!sending) {
      setSendingStage(0);
      setSendingAddress('');
      setInvalidAddress(false);
      setCoinToSend();
      setCoinToSendName();
      setCoinToSendBalance();
      setSendingUserOp();
      setSendingAmount();
      setLoading(false);
      setSent();
    }
  }, [sending]);



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

  const getWalletBalances = async () => {

    let USDC = await getERC20Balance(web3, addresses.usdc, walletAddress)
      .catch(() => console.error('Couldn\'t get USDC balance'));
    let SFLUV = await getERC20Balance(web3, addresses.sfluv, walletAddress)
      .catch(() => console.error('Couldn\'t get SFLUV balance'));
    let MATIC = await getERC20Balance(web3, addresses.matic, walletAddress)
      .catch(() => console.error('Couldn\'t get MATIC balance'));
    let USDCe = await getERC20Balance(web3, addresses.usdce, walletAddress)
      .catch(() => console.error('Couldn\'t get USDC.e balance'));

    setUSDCBalance(formatBalance(USDC));
    setUSDCeBalance(formatBalance(USDCe));
    setSFLUVBalance(formatBalance(SFLUV));
    setMATICBalance(formatBalance(MATIC / 1000000000000));

  }


  useEffect(() => {
    if(walletAddress) {
      getWalletBalances();
    }
  }, [walletAddress]);

  useEffect(() => {

    if(web3) {
      const getWalletAddress = async () => {
        console.log(web3);
        const owner = (await web3.eth.getAccounts())[0];
        const wallet = await axios.get(`${links.server}/erc4337/sender-info?owner=${owner}`)
          .then((res) => JSON.parse(res.data).sender)
          .catch(console.error)
        console.log('Owner: ' + owner);
        console.log('Wallet: ' + wallet);
        setOwnerAccountAddress(owner);
        setWalletAddress(wallet);
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

  const toggleSendModal = () => {
    setSending(!sending);
  }

  Modal.setAppElement('#app');

  const formatHash = (hash) => {
    if(hash.slice(0, 2) === '0x') {
      hash = hash.slice(2);
    }

    return hash;
  }

  const makeSendContract = async (tokenContractAddress, amount, toAddress = walletAddress, owner = ownerAccountAddress) => {


    const transferLink = `${links.server}/erc4337/userop/transfer?target=${formatHash(tokenContractAddress)}&to=${formatHash(toAddress)}&amount=${amount}&owner=${formatHash(owner)}
    `
    const contract = await axios.get(transferLink)
      .catch((err) => console.error(err));
    console.log(contract);
    return contract;
  }

  const executeSend = async (userOpData) => {

    const privateKey = await getPrivateKey(web3auth);

    const signedUserOp = await signUserOp(web3, privateKey, userOpData, ownerAccountAddress);

    await axios.post(
      `${links.server}/erc4337/userop/send`,
      {entryPoint: process.env.REACT_APP_ENTRYPOINT, op: signedUserOp},
      {headers: {
        'Content-Type': 'application/json',
        'Accept': 'application/json'
      }}
    )
  }

  const checkStage0 = (coinAddress, name, balance) => {
    setCoinToSend(coinAddress);
    setCoinToSendName(name);
    setCoinToSendBalance(balance);
    setSendingStage(1);
  }

  const checkStage1 = (address = sendingAddress) => {
    if(web3.utils.isAddress(formatHash(address))) {
      setSendingAddress(formatHash(address))
      setInvalidAddress(false);
      setSendingStage(2);
    } else {
      setInvalidAddress(true);
    }
  }

  const checkStage2 = async (amount) => {

    const formattedSendAmount = amount * 1000000

    const userOpData = await makeSendContract(coinToSend, formattedSendAmount, sendingAddress, ownerAccountAddress)
      .then((res) => res.data)
      .catch(console.error);

    setSendingUserOp(userOpData);
    setSendingStage(3);
  }

  const sendMATIC = () => {

  }

  useEffect(() => {

    if(sendingAmount?.toString().split('.')[1]?.length > 2) {
      setSendingAmount(formatBalance(sendingAmount));
    }

    if(Number(sendingAmount) > Number(coinToSendBalance)) {
      setTooHigh(true);
    } else {
      setTooHigh(false);
    }

    if(Number(sendingAmount) === 0) {
      setDisableSubmit(true);
    } else {
      setDisableSubmit(false);
    }

    if(Number(sendingAmount) < 0 || sendingAmount === undefined) {
      setSendingAmount(0);
    }

  }, [sendingAmount])

  return (
    <>
      <div id='profilePage' className='pageTemplate'>
        {walletAddress && <><div id='profileSections'>
          <div className='profileSection'>
            <span className='col1'>SFLUV: </span>
            <span className='col2'>{SFLUVBalance}</span>
          </div>
          <div className='profileSection'>
            <span className='col1'>USDC: </span>
            <span className='col2'>{USDCBalance}</span>
          </div>
          <div className='profileSection'>
            <span className='col1'>USDC.e: </span>
            <span className='col2'>{USDCeBalance}</span>
          </div>
          <div  className='profileSection capitalized'>
            <span className='col1'>MATIC: </span>
            <span className='col2'>{MATICBalance}</span>
          </div>
        </div>
        {walletAddress && (
          <>
            <button
              className='button1 walletButton'
              onClick={toggleDepositModal}
            >
              Deposit
            </button>
            <button
              className='button2 walletButton'
              onClick={toggleSendModal}
            >
              Send
            </button>
          </>
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
        <Modal
          id='sendModal'
          isOpen={sending}
          onRequestClose={toggleSendModal}
          className='pageTemplate'
          contentElement={
            (props, contentElement) => (
              <div {...props}>
                {loading ?
                  <div style={{paddingTop: '20vh', paddingBottom: '20vh'}}>
                    <CircleLoader color='#eb6c6c' loading={loading}/>
                  </div>
                  :
                  <>
                  {sendingStage === 0 && <div>
                    <h1>Choose Coin</h1>
                    {coins.map((coin) => (
                      <div
                        className='sendingCoinItem'
                        onClick={() => checkStage0(coin.address, coin.name, coin.balance)}
                        key={coin.name}
                      >
                        <div>
                          <p><b>{coin.name}</b></p>
                        </div>
                        <div>
                          <p>{coin.balance}</p>
                        </div>
                      </div>
                    ))}
                    <p id='warningMessage'><b>Please</b> be sure all transactions are being sent on <b>{process.env.REACT_APP_CHAIN_DISPLAY_NAME}</b></p>
                  </div>}
                  {sendingStage === 1 && <div>
                    <h1 style={{marginBottom: '0px'}}>To Address</h1>
                    <CircleLoader color='#eb6c6c' loading={true} id='qrLoader'/>
                    <QrReader
                      onResult={(result, error) => {
                        if(result?.text) {
                          checkStage1(result.text);
                        }
                      }}
                      containerStyle={{
                        width: '25vh',
                        height: '25vh',
                        margin: 'auto',
                        overflow: 'hidden',
                        borderRadius: '40px'
                      }}
                      videoContainerStyle={{
                        transform: 'scale(1.5)',
                      }}
                      scanDelay={500}
                    >
                    </QrReader>
                    <form onSubmit={(e) => {
                      e.preventDefault();
                      checkStage1();
                    }}>
                      <label id='addressLabel'>Address:</label>
                      <input id='recipientAddressInput'
                        type='text'
                        onChange={(e) => {
                          e.preventDefault();
                          setSendingAddress(e.target.value);
                        }}
                        value={sendingAddress}
                        placeholder={walletAddress}
                      />
                      <button
                        className='button2'
                        id='nextButton'
                        type='submit'
                      >
                        Next
                      </button>
                      {invalidAddress && <p>Input must be a valid blockchain address</p>}
                    </form>
                  </div>}
                  {sendingStage === 2 && <div>
                    <h1>Amount</h1>
                    <h3>{coinToSendName}</h3>
                    <p>Balance: {coinToSendBalance}</p>
                    <form onSubmit={async (e) => {
                      e.preventDefault();
                      setLoading(true);
                      await checkStage2(sendingAmount)
                        .catch(console.error);
                      setLoading(false);
                    }}>
                      <input
                        id='sendingAmount'
                        value={sendingAmount}
                        onChange={(e) => {
                          setSendingAmount(e.target.value);
                        }}
                        placeholder={0}
                        type='number'
                        step='0.01'
                      />
                      {tooHigh && <p>Not enough balance</p>}
                      <button
                        className={disableSubmit || tooHigh ? 'disabledButton' : 'button2'}
                        id='nextButton'
                        disabled={disableSubmit || tooHigh}
                      >
                        Next
                      </button>
                    </form>
                  </div>}
                  {sendingStage === 3 && <div>
                    <h1>Confirmation</h1>
                    <p>Double check the details</p>
                    <div id='profileSections'>
                      <div className='profileSection'>
                        <span className='col1'><b>Coin: </b></span>
                        <span className='col2'>{coinToSendName}</span>
                      </div>
                      <div className='profileSection'>
                        <span className='col1'><b>To: </b></span>
                        <span className='col2' id='sendingAddress'>{sendingAddress}</span>
                      </div>
                      <div className='profileSection'>
                        <span className='col1'><b>Amount: </b></span>
                        <span className='col2'>{sendingAmount}</span>
                      </div>
                    </div>
                    <button
                      onClick={async () => {
                        setLoading(true);
                        await executeSend(sendingUserOp)
                          .catch(console.error)
                        setLoading(false);
                        setSendingStage(4);
                        setTimeout(() => {
                          setSending(false);
                        }, 5000)
                      }}
                      className='button2'
                      id='nextButton'
                    >
                      Send
                    </button>
                  </div>}
                  {sendingStage === 4 && <div>
                    <h1>Sent!</h1>
                    <AiFillCheckCircle style={{color: '#eb6c6c'}} size='7vh'/>
                    <h3>{coinToSendName}</h3>
                    <div id='profileSections'>
                      <div className='profileSection'>
                        <span className='col1'><b>To: </b></span>
                        <span className='col2' id='sendingAddress'>{sendingAddress}</span>
                      </div>
                      <div className='profileSection'>
                        <span className='col1'><b>Amount: </b></span>
                        <span className='col2'>{sendingAmount}</span>
                      </div>
                      <p style={{maxWidth: '90%', margin: 'auto', marginTop: '5px'}}>Balance in wallet may take up to a minute to update</p>
                    </div>
                  </div>}
                  </>
                }
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
        </Modal></>}
        {!walletAddress && <div style={{paddingTop: '20vh', paddingBottom: '20vh'}}>
          <CircleLoader color='#eb6c6c' loading={true}/>
        </div>}
      </div>
    </>
  )
}

export default Wallet;