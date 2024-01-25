import { EthereumPrivateKeyProvider } from "@web3auth/ethereum-provider";
import { CHAIN_NAMESPACES } from "@web3auth/base";
import { Web3AuthNoModal } from "@web3auth/no-modal";
import { OpenloginAdapter } from "@web3auth/openlogin-adapter";
import Logo from '../assets/SFLUV LOGO-1.svg';
import 'dotenv';

const clientId = process.env.REACT_APP_CLIENT_ID;

const chainConfig = {
  chainNamespace: CHAIN_NAMESPACES.EIP155,
  chainId: process.env.REACT_APP_CHAIN_ID,
  rpcTarget: "https://rpc.ankr.com/polygon",
  displayName: process.env.REACT_APP_CHAIN_DISPLAY_NAME,
  blockExplorer: "https://polygon.etherscan.io",
  ticker: "MATIC",
  tickerName: "Polygon",
};

const privateKeyProvider = new EthereumPrivateKeyProvider({
  config: { chainConfig },
});

const openloginAdapter = new OpenloginAdapter({
  privateKeyProvider: privateKeyProvider,
  adapterSettings: {
    uxMode: "redirect",
    whiteLabel: {
      appName: "SFLUV",
      appUrl: "https://web3auth.io",
      logoLight: Logo,
      logoDark: Logo,
      defaultLanguage: "en", // en, de, ja, ko, zh, es, fr, pt, nl
      mode: "light", // whether to enable dark mode. defaultValue: auto
      theme: {
        primary: "#whitesmoke",
      },
      useLogoLoader: true,
    },
  }
});

const web3auth = new Web3AuthNoModal({
  clientId,
  chainConfig,
  web3authNetwork: process.env.REACT_APP_AUTH_NETWORK,
});

web3auth.configureAdapter(openloginAdapter);

export default web3auth;