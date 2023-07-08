import { useEffect, useState } from "react";
import { useWeb3 } from "./context/Web3Context";
import { useUser } from "./context/UserContext";

const useBalance = () => {
    const { user } = useUser();
    const { web3 } = useWeb3();
    const [balance, setBalance] = useState("...")
    useEffect(() => {
        const getBalance = async () => {
          if (!user || !web3) return
          try {
            // If account and web3 are available, get the balance
            const balance = await web3.eth.getBalance(user)
    
            // Convert the balance from Wei to Ether and set the state variable
            setBalance(web3.utils.fromWei(balance).substring(0, 7))
          } catch (error) {
            console.error(error)
          }
        }
    
        getBalance()
      }, [user])
    return (balance);
}
 
export default useBalance;