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
  sendOtp: ( otp : string ) => void,
  sendPhone: ( phoneNumber : string ) => void,
  reSendPhone: () => void,
  initializeRatio: () => void
}

// Create a context for user data.
const RatioContext = createContext<RatioContextType>({
  bearer: null,
  ratio: null,
  phoneId: null,
  sendOtp: ( otp: string ) => {},
  sendPhone: ( phoneNumber : string ) => {},
  reSendPhone: () => {},
  initializeRatio: () => {}
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

    
    const [bearer, setBearer] = useState('');

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

    const [phoneId, setPhoneId] = useState('');

    const initializeRatio = async () => {
      if(user && web3 && !ratio) {
        await ratioFirstChallenge( luv_server + '/ratio/wallet' )
        .then((res) => {
          setBearer(JSON.parse(res).jwt)
        })
        .catch((e) => {
            console.log(e.message);
      })}
    }

    const [tempPhone, setTempPhone] = useState('');

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



    // Update the user state with the first account (if available), otherwise set to null.
    
  



  return (
    <RatioContext.Provider
      value={{
        bearer: bearer,
        ratio: ratio,
        phoneId: phoneId,
        sendPhone,
        reSendPhone,
        sendOtp,
        initializeRatio
      }}
    >
      {children}
    </RatioContext.Provider>
  )
}
