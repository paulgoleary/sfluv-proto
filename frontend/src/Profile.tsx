import { VStack } from "@chakra-ui/react"
import { useUser } from "./context/UserContext"
import ConnectButton from "./components/ConnectButton"
import WalletDetail from "./components/WalletDetail"
import DisconnectButton from "./components/DisconnectButton"
import ShowUIButton from "./components/ShowUIButton"
import SignMessage from "./components/SignMessage"
import { useRatio } from "./context/RatioContext"
import { Button } from "@chakra-ui/react"
import { useEffect } from "react"
import { useWeb3 } from "./context/Web3Context"
import { magic } from "./libs/magic"

const Profile = () => {
  const { user } = useUser();
  const ratio = useRatio();
  const qrSource = 'https://api.qrserver.com/v1/create-qr-code/?data=ethereum:'+ user +'&size=100x100'
  const { initializeWeb3 } = useWeb3()


  // Define the event handler for the button click
  const handleConnect = async () => {
    try {
      // Try to connect to the wallet using Magic's user interface
      await magic.wallet.connectWithUI()

      // If connection to the wallet was successful, initialize new Web3 instance
      initializeWeb3()
    } catch (error) {
      // Log any errors that occur during the connection process
      console.error("handleConnect:", error)
    }
  }
  useEffect(() => {
    if(!user){
      handleConnect();
    }
  },[])


  return ( 
      <VStack className="magicUI" justifyContent="center" alignItems="center" minH="30vh" maxH="100vh">
      {!user ? (<Button className="button" onClick={() => {handleConnect()}}>Log In</Button>
        ) : (
          <>
            {ratio.ratio && <h1 className="loggedIn">Logged in with Ratio</h1>}
            <img src={qrSource} alt="AddressQR"/>
            <WalletDetail />
            <SignMessage />
            <ShowUIButton />
            <DisconnectButton /> 
          </>
        )}
      </VStack>
    );
}
 
export default Profile;