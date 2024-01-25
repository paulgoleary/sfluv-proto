import React, { useState, useContext, useEffect } from 'react';
import CircleLoader from 'react-spinners/ClipLoader';
import Web3Context from '../Web3Context.js';
import getERC20Balance from '../models/getERC20Balance';
import getPrivateKey from '../models/getPrivateKey.js';
import signUserOp from '../models/signUserOp.js';
import addresses from '../../utils/addresses.js';
import links from '../../utils/links.js';
import axios from 'axios';
import '../styles/UnwrapTool.css';
import 'dotenv';

const UnwrapTool = () => {

  const [walletAddress, setWalletAddress] = useState();
  const [ownerAccountAddress, setOwnerAccountAddress] = useState();
  const [SFLUVBalance, setSFLUVBalance] = useState();
  const [unwrapAmount, setUnwrapAmount] = useState(0.00);
  const [tooHigh, setTooHigh] = useState(false);
  const [disableSubmit, setDisableSubmit] = useState(false);

  const { web3, web3auth } = useContext(Web3Context);

  useEffect(() => {
    if(web3) {
      getAddresses();
    }
  }, [web3]);

  const getSFLUVBalance = async (wallet) => {
    const balance = await getERC20Balance(web3, addresses.sfluv, wallet);
    setSFLUVBalance(formatBalance(balance));
  }

  const getAddresses = async () => {
    const owner = (await web3.eth.getAccounts())[0];
    const wallet = await axios.get(`${links.server}/erc4337/sender-info?owner=${owner}`)
      .then((res) => JSON.parse(res.data).sender)
      .catch(console.error);
    await getSFLUVBalance(wallet);

    setOwnerAccountAddress(owner);
    setWalletAddress(wallet);
  }

  const formatBalance = (balance) => {
    balance = (balance * 1.0).toFixed(2);

    return balance
  }


  const makeUnwrapContract = async (tokenContractAddress, amount, toAddress = walletAddress, owner = ownerAccountAddress) => {

    console.log('Owner: ' + owner);
    console.log('To: ' + toAddress);
    console.log('Amount: ' + amount);
    console.log('Contract Address: ' + tokenContractAddress);

    const formatHash = (hash) => {
      console.log('PREFORMAT ' + hash);
      if(hash.slice(0, 2) === '0x') {
        hash = hash.slice(2);
      }

      console.log('POSTFORMAT ' + hash);
      return hash;
    }

    const transferLink = `${links.server}/erc4337/userop/withdrawto?target=${formatHash(tokenContractAddress)}&to=${toAddress}&amount=${amount}&owner=${owner}
    `
    const contract = await axios.get(transferLink)
      .catch((err) => console.error(err));
    return contract;
  }

  const executeUnwrap = async () => {

    const formattedUnwrapAmount = unwrapAmount * 1000000

    const userOpData = await makeUnwrapContract(
      addresses.sfluv,
      formattedUnwrapAmount,
      walletAddress,
      ownerAccountAddress
    )
      .then((res) => res.data)
      .catch(console.error);

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
      .then((res) => {
        getSFLUVBalance(walletAddress);
      })

  }

  useEffect(() => {

    if(unwrapAmount?.toString().split('.')[1]?.length > 2) {
      setUnwrapAmount(formatBalance(unwrapAmount));
    }

    if(Number(unwrapAmount) > Number(SFLUVBalance)) {
      setTooHigh(true);
    } else {
      setTooHigh(false);
    }

    if(Number(unwrapAmount) === 0) {
      setDisableSubmit(true);
    } else {
      setDisableSubmit(false);
    }

    if(Number(unwrapAmount) < 0 || unwrapAmount === undefined) {
      setUnwrapAmount(0);
    }

  }, [unwrapAmount])

  return (
    <>
      <div id='unwrapPage' className='pageTemplate'>
        {walletAddress && <><h1>Unwrap Tokens</h1>
        <form onSubmit={(e) => {
          e.preventDefault();
          executeUnwrap();
        }}>
          <input
            id='unwrapAmount'
            type='number'
            step='0.01'
            value={unwrapAmount}
            onChange={(e) => setUnwrapAmount(e.target.value)}
          />
          {tooHigh && <p style={{color: '#eb6c6c'}}>Not enough balance</p>}
          <p>Available balance: {SFLUVBalance && formatBalance(SFLUVBalance)}</p>
          <button
            type='submit'
            disabled={tooHigh || disableSubmit}
            className={`${tooHigh || disableSubmit ? 'disabledButton' : 'button2'} unwrapButton`}
          >
            Submit
          </button>
        </form></>}
        {!walletAddress && <div style={{paddingTop: '20vh', paddingBottom: '20vh'}}>
          <CircleLoader color='#eb6c6c' loading={true}/>
        </div>}
      </div>
    </>
  )
}

export default UnwrapTool;