import React, { createContext, useContext, useEffect, useState } from "react"
import { useWeb3 } from "./Web3Context"
import { useUser } from "./UserContext"
import links from "../links"
import { Button } from "@chakra-ui/react"

// Define the type for the user context.
type RatioContextType = {
  ratio: string | null,
  initializeRatio: ( phoneNumber: string ) => void
}

// Create a context for user data.
const RatioContext = createContext<RatioContextType>({
  ratio: null,
  initializeRatio: ( phoneNumber: string ) => {},
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

    
    

    const ratioLogin = async ( link:string, phoneNumber:string ) => {
      
      const resp = JSON.parse(await ratioFirstChallenge( luv_server + '/ratio/wallet' )).challenge;
      console.log(resp);
      const smsInputs = {
        method: 'POST',
        headers: {'Content-Type': 'application/json', 'Accept': 'application/json', 'Authorization': 'Bearer ' + resp},
        body: JSON.stringify({phoneNumber})
      }
      const response = await fetch(link, smsInputs);
      if (!response.ok) {
          throw new Error('Data could not be fetched!')
      } else {
        console.log(response);
        return response.json();
      }}


    const initializeRatio = async ( phoneNumber:string ) => {
      if(user && web3 && !ratio) {
        await ratioLogin( luv_server + '/ratio/auth/sms', phoneNumber )
        .then((res) => {
            console.log(res);
            setRatio(JSON.parse(res));
        })
        .catch((e) => {
            console.log(e.message);
      })}
    }



    // Update the user state with the first account (if available), otherwise set to null.
    
  



  return (
    <RatioContext.Provider
      value={{
        ratio: ratio,
        initializeRatio
      }}
    >
      {children}
    </RatioContext.Provider>
  )
}
