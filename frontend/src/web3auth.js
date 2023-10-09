import { Web3Auth } from "@web3auth/modal";

// Initialize within useEffect()
const web3auth = new Web3Auth({
  clientId: "BJfhyvUB8U_a8MwfgBI25Fjv2DY9tT7MGxeEvzRef_Q8CUJ30yVX3MDZUl0UDyvFkbaKxIr-6xZRpNz43Hdh6cI", // Get your Client ID from the Web3Auth Dashboard
  web3AuthNetwork: "sapphire_mainnet", // Web3Auth Network
  chainConfig: {
    chainNamespace: "eip155",
    chainId: "0x89",
    rpcTarget: "https://rpc.ankr.com/polygon",
    displayName: "Polygon Mainnet",
    blockExplorer: "https://polygon.etherscan.io",
    ticker: "MATIC",
    tickerName: "Polygon",
  }
});


const auth = { web3auth }



export default auth;