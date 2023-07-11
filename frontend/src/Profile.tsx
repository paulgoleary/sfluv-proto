import { VStack } from "@chakra-ui/react"
import { useUser } from "./context/UserContext"
import { useWeb3 } from "./context/Web3Context"
import ConnectButton from "./components/ConnectButton"
import WalletDetail from "./components/WalletDetail"
import DisconnectButton from "./components/DisconnectButton"
import ShowUIButton from "./components/ShowUIButton"
import SignMessage from "./components/SignMessage"
import BuyButton from "./components/BuyButton"
import GetFetchData from "./GetFetchData"
import links from "./links"
import { useEffect, useState } from "react"

const Profile = () => {
  const { web3 } = useWeb3();
  const { user } = useUser()
  const { luv_server } = links();
  const walletAddress = useUser().user;
  const walletNetwork = "POLYGON";
  const [preSignature, setPreSignature] = useState('');
  const inputs = { walletAddress, walletNetwork };
  const inputData = {
    method: 'POST',
    headers: {'Content-Type': 'application/json', 'Accept': 'application/json'},
    body: JSON.stringify(inputs)
  }
  const firstFetch = async ( link:string, postData:object ) => {
    console.log(inputs);
    const response = await fetch(link, postData);
    console.log(inputs);
    if (!response.ok) {
        throw new Error('Data could not be fetched!')
      } else {
        return response.json()
      }

  }

  const [signedData, setSignedData] = useState('');

  const MagicSign = async ( message:string ) => {
    var signedMessage;
    if (user && web3 && message != '') {
      try {
        signedMessage = await web3.eth.personal.sign(message, user, "")
        setSignedData(signedMessage)
        console.log(signedData);
      } catch (error) {
        console.error("handleSignMessage:", error)
      }
    }
    return signedMessage;
  }

  const [parsed, setParsed] = useState('');
  const ratioLogin = async ( link:string ) => {
    const challengeData = await firstFetch( luv_server + '/ratio/sessions', inputData);
    console.log('Challenge: ' + challengeData);
    const parsedData = JSON.parse(challengeData).challenge;
    setParsed(parsedData);
    console.log('Parsed: ' + parsedData);
    const signature = await MagicSign(parsedData);
    console.log(signature);
    const signedInputs = { walletAddress, walletNetwork, signature };
    console.log(signedInputs);
    const signedChallengeInputs = {
      method: 'POST',
      headers: {'Content-Type': 'application/json', 'Accept': 'application/json'},
      body: JSON.stringify(signedInputs)
    }
    const response = await fetch(link, signedChallengeInputs);
    console.log(signedChallengeInputs);
    if (!response.ok) {
        throw new Error('Data could not be fetched!')
      } else {
        console.log(response);
        return response.json();
      }

  } 

  const [data, setData] = useState();

  useEffect(() => {
    if(user && web3) {
      ratioLogin( luv_server + '/ratio/wallet' )
      .then((res) => {;
          setData(res);
      })
      .catch((e) => {
          console.log(e.message);
    })}
  }, [user]); 

  
  



  return ( 
      <VStack className="magicUI" justifyContent="center" alignItems="center" minH="30vh" maxH="100vh">
      {!user ? (
          <ConnectButton />
        ) : (
          <>
            <WalletDetail />
            <SignMessage />
            <ShowUIButton />
            <DisconnectButton />
          </>
        )}
        {user && (<button>Ratio</button>)}
        <div>
          <h2>Parsed: {parsed}</h2>
          <br/>
          <h2>Signed: {signedData}</h2>
          <br/>
          <h2>Response: {data}</h2>
        </div>
      </VStack>
    );
}
 
export default Profile;