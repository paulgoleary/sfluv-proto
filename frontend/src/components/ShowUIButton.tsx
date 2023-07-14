import { useEffect, useState } from "react"
import { Button, Spinner } from "@chakra-ui/react"
import { magic } from "../libs/magic"

const ShowUIButton = () => {
  // Initialize state variable to decide whether to show button or not
  const [showButton, setShowButton] = useState(false);

  // Define a function to check the type of the wallet
  const checkWalletType = async () => {
    try {
      // Fetch the wallet's information using Magic's user.getInfo method
      const walletInfo = await magic.user.getInfo()
      console.log(walletInfo)

      ///@ts-ignore
      // Determine if the wallet type is "magic"
      const isMagicWallet = walletInfo.walletType === "magic"

      // Set 'showButton' state based on the result of the check
      setShowButton(isMagicWallet)
    } catch (error) {
      // Log any errors that occur during the wallet type check process
      console.error("checkWalletType:", error)
    }
  }

  useEffect(() => {
    // Call the checkWalletType function
    checkWalletType()
  }, [magic])

  // Define the event handler for the button click
  const handleShowUI = async () => {
    try {
      // Try to show the magic wallet user interface
      await magic?.wallet.showUI()
    } catch (error) {
      // Log any errors that occur during the process
      console.error("handleShowUI:", error)
    }
  }
  return (
    <div>
      {showButton && <Button className="button" onClick={handleShowUI}>Show UI</Button>}
      {!showButton && <Spinner/>}
    </div> 
  ) 
}

  // Render the button component if showButton is true, otherwise render nothing
  

  

export default ShowUIButton
