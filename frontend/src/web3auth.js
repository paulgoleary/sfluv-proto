import { Web3Auth } from "@web3auth/modal";

// Initialize within useEffect()
const web3auth = new Web3Auth({
  clientId: "BG6JbfWeGCYLBd3SXfOjrMTzKN8Bpk2I4cUh2rqoSWNelWMC5qSdzyfe-ecnTJS55MbO5epo84X0pQdns8m9J8U", // Get your Client ID from the Web3Auth Dashboard
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
