import { VStack } from "@chakra-ui/react"
import { useUser } from "./context/UserContext"
import ConnectButton from "./components/ConnectButton"
import WalletDetail from "./components/WalletDetail"
import DisconnectButton from "./components/DisconnectButton"
import ShowUIButton from "./components/ShowUIButton"
import SignMessage from "./components/SignMessage"
import BuyButton from "./components/BuyButton";
import Profile from "./Profile"
import Navbar from "./Navbar"
import Home from "./Home"
import About from "./About"
import Test from "./Test"
import Buy from "./Buy"
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import { useEffect, useState } from "react"
import { useWeb3 } from "./context/Web3Context"
import links from "./links"
import { useRatio } from "./context/RatioContext"


function App() {
  const { web3 } = useWeb3();
  const { user } = useUser();
  const { userId } = useRatio();
  // const { luv_server } = links();
  // const walletAddress = useUser().user;
  // const walletNetwork = "POLYGON";
  // const inputs = { walletAddress, walletNetwork };
  // const [ratioString, setRatioString] = useState();
  // const inputData = {
  //   method: 'POST',
  //   headers: {'Content-Type': 'application/json', 'Accept': 'application/json'},
  //   body: JSON.stringify(inputs)
  // }
  // const firstFetch = async ( link:string, postData:object ) => {
  //   console.log(inputs);
  //   const response = await fetch(link, postData);
  //   console.log(inputs);
  //   if (!response.ok) {
  //       throw new Error('Data could not be fetched!')
  //     } else {
  //       return response.json()
  //     }

  // }

  // const [signedData, setSignedData] = useState('');

  // const MagicSign = async ( message:string ) => {
  //   var signedMessage;
  //   if (user && web3 && message != '') {
  //     try {
  //       signedMessage = await web3.eth.personal.sign(message, user, "")
  //       setSignedData(signedMessage)
  //       console.log(signedData);
  //     } catch (error) {
  //       console.error("handleSignMessage:", error)
  //     }
  //   }
  //   return signedMessage;
  // }

 
  // const ratioLogin = async ( link:string ) => {
  //   const challengeData = await firstFetch( luv_server + '/ratio/sessions', inputData);
  //   console.log('Challenge: ' + challengeData);
  //   const parsedData = JSON.parse(challengeData).challenge;
  //   console.log('Parsed: ' + parsedData);
  //   const signature = await MagicSign(parsedData);
  //   console.log(signature);
  //   const signedInputs = { walletAddress, walletNetwork, signature };
  //   console.log(signedInputs);
  //   const signedChallengeInputs = {
  //     method: 'POST',
  //     headers: {'Content-Type': 'application/json', 'Accept': 'application/json'},
  //     body: JSON.stringify(signedInputs)
  //   }
  //   const response = await fetch(link, signedChallengeInputs);
  //   console.log(signedChallengeInputs);
  //   if (!response.ok) {
  //       throw new Error('Data could not be fetched!')
  //     } else {
        
  //       return response.json();
  //     }

  // } 

  // const [data, setData] = useState();

  // useEffect(() => {
  //   if(user && web3 && !ratioString) {
  //     ratioLogin( luv_server + '/ratio/wallet' )
  //     .then((res) => {;
  //         console.log(res);
  //         setData(res);
  //         setRatioString(JSON.parse(res).challenge);
  //     })
  //     .catch((e) => {
  //         console.log(e.message);
  //   })}
  // }, [user]);


  return (
      <Router>
        <div className="App">
          <Navbar/>
              <Routes>
                <Route path="/profile" element={<Profile/>}/>
                <Route path="/" element={<Home/>}/>
                <Route path="/about" element={<About/>}/>
                <Route path="/test" element={<Test/>}/>
                {user && <Route path="/buy" element={<Buy/>}/>}
              </Routes>
        </div>
      </Router>
  )
}


export default App
