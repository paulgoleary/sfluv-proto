import { VStack } from "@chakra-ui/react"
import { useUser } from "./context/UserContext"
import ConnectButton from "./components/ConnectButton"
import WalletDetail from "./components/WalletDetail"
import DisconnectButton from "./components/DisconnectButton"
import ShowUIButton from "./components/ShowUIButton"
import SignMessage from "./components/SignMessage"
import BuyButton from "./components/BuyButton";

const Profile = () => {
    const { user } = useUser()
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
        </VStack>
     );
}
 
export default Profile;