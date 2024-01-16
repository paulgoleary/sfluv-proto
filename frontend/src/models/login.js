import web3auth from '../web3auth.js';
import { useContext } from 'react';

const login = async () => {
  const web3authProvider = await web3auth.connectTo(
    WALLET_ADAPTERS.OPENLOGIN,
    {
      loginProvider: "google",
    }
  );
  console.log(web3authProvider);
  setProvider(web3authProvider);
  if (web3auth.connected) {
    setLoggedIn(true);
  }
};