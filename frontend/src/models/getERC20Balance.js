import { useContext } from 'react';
import Web3Context from '../Web3Context.js';


const getERC20Balance = async (web3, contractAddress, walletAddress) => {


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

export default getERC20Balance;