import React, { createContext, useContext, useEffect, useState } from "react"
import { useWeb3 } from "./Web3Context"
import { useUser } from "./UserContext"
import links from "../links"
import { Button } from "@chakra-ui/react"

// Define the type for the user context.
type RatioContextType = {
  bearer: string | null,
  ratio: string | null,
  phoneId: string | null,
  triedLogin: boolean | null,
  freezeUser: string | null,
  sendOtp: ( otp : string ) => void,
  sendPhone: ( phoneNumber : string ) => void,
  reSendPhone: () => void,
  initializeRatio: () => void,
  resetTriedLogin: () => void,
  resetRatioState: () => void
}

// Create a context for user data.
const RatioContext = createContext<RatioContextType>({
  bearer: null,
  ratio: null,
  phoneId: null,
  triedLogin: null,
  freezeUser: null,
  sendOtp: ( otp: string ) => {},
  sendPhone: ( phoneNumber : string ) => {},
  reSendPhone: () => {},
  initializeRatio: () => {},
  resetTriedLogin: () => {},
  resetRatioState: () => {}
})

// Custom hook for accessing user context data.
export const useRatio = () => useContext(RatioContext)

// Provider component that wraps parts of the app that need user context.
export const RatioProvider = ({ children }: { children: React.ReactNode }) => {
  // Use the web3 context.
  const { web3 } = useWeb3()
  const { user } = useUser()
  const [ratio, setRatio] = useState<string | null>(null)
  const { luv_server } = links();
  const walletAddress = useUser().user;
  const walletNetwork = "POLYGON";
  const inputs = { walletAddress, walletNetwork };
  const inputData = {
    method: 'POST',
    headers: {'Content-Type': 'application/json', 'Accept': 'application/json'},
    body: JSON.stringify(inputs)
  }
  const [signedData, setSignedData] = useState('');
  const [bearer, setBearer] = useState('');
  const [phoneId, setPhoneId] = useState('');
  const [triedLogin, setTriedLogin] = useState(false);
  const [tempPhone, setTempPhone] = useState('');
  const [freezeUser, setFreezeUser] = useState('');






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

 
  const ratioFirstChallenge = async ( link:string ) => {
    const challengeData = await firstFetch( luv_server + '/ratio/sessions', inputData);
    console.log('Challenge: ' + challengeData);
    const parsedData = JSON.parse(challengeData).challenge;
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

    
  

  const ratioLogin = async ( link:string, phoneNumber:string ) => {
    const smsInputs = {
      method: 'POST',
      headers: {'Content-Type': 'application/json', 'Accept': 'application/json', 'Authorization': 'Bearer ' + bearer},
      body: JSON.stringify({phoneNumber})
    }
    console.log(smsInputs);
    const response = await fetch(link, smsInputs);
    if (!response.ok) {
        throw new Error('Data could not be fetched!')
    } else {
      console.log(response);
      return response.json();
  }}

  

  const initializeRatio = async () => {
    if(user && web3 && !ratio) {
      setFreezeUser(user);
      await ratioFirstChallenge( luv_server + '/ratio/wallet' )
      .then((res) => {
        setBearer(JSON.parse(res).jwt)
      })
      .catch(() => {
          console.log('Unable to initialize Ratio login. Please try again.');
          setTriedLogin(true);
    })}
  }


  const sendPhone = async ( phoneNumber: string ) => {
    setTempPhone(phoneNumber);
    if(user && web3 && !ratio) {
      await ratioLogin( luv_server + '/ratio/jwt/sms-send', phoneNumber )
      .then((res) => {
          console.log(JSON.parse(res).phoneId);
          
          setPhoneId(JSON.parse(res).phoneId);
      })
      .catch((e) => {
          console.log(e.message);
    })}
  }

  const reSendPhone = async () => {
    setTempPhone(tempPhone);
    if(user && web3 && !ratio) {
      await ratioLogin( luv_server + '/ratio/jwt/sms-send', tempPhone )
      .then((res) => {
          console.log(JSON.parse(res).phoneId);
          
          setPhoneId(JSON.parse(res).phoneId);
      })
      .catch((e) => {
          console.log(e.message);
    })}
  }

  const makeOtpSend = async ( otp : string ) => {
    console.log(phoneId);
    const response = await fetch(luv_server + '/ratio/jwt/sms-auth', {
      method: 'POST',
      headers: {'Content-Type': 'application/json', 'Accept': 'application/json', 'Authorization': 'Bearer ' + bearer},
      body: JSON.stringify({phoneId, otp})
    })
    if (!response.ok) {
      throw new Error('Data could not be fetched!')
    } else {
    console.log(response);
    return response.json();
    } 
  }

  const sendOtp = async ( otp: string ) => {
    if(user && web3 && !ratio) {
      await makeOtpSend(otp)
      .then((res) => {
          console.log(res);
          setRatio(res);
      })
      .catch((e) => {
          console.log(e.message);
      })
    } 
  }

  const resetTriedLogin = () => {
    if(triedLogin) {
      setTriedLogin(false);
    }
  }  

  const resetRatioState = () => {
    setSignedData('');
    setBearer('');
    setPhoneId('');
    setTriedLogin(false);
    setTempPhone('');
    setRatio('');
  }



  return (
    <RatioContext.Provider
      value={{
        bearer: bearer,
        ratio: ratio,
        phoneId: phoneId,
        triedLogin: triedLogin,
        freezeUser: freezeUser,
        sendPhone,
        reSendPhone,
        sendOtp,
        initializeRatio,
        resetTriedLogin,
        resetRatioState
      }}
    >
      {children}
    </RatioContext.Provider>
  )
}
