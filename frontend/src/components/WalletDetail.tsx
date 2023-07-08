import { useEffect, useState } from "react"
import { Text, Flex, Divider } from "@chakra-ui/react"
import { useWeb3 } from "../context/Web3Context"
import { useUser } from "../context/UserContext"
import useBalance from "../useBalance"

const WalletDetail = () => {
  // Use the Web3Context to get the current instance of web3
  const { web3 } = useWeb3()
  // Use the UserContext to get the current logged-in user
  const { user } = useUser()
  const balance = useBalance();

  // Render the account address and balance
  return (
    <Flex className="walletDetails" direction="column">
      <Text fontWeight="bold">Address</Text>
      <Text my={2}>
        {user}
      </Text>
      <Divider my={2} />
      <Text fontWeight="bold">Balance</Text>
      <Text>{balance} ETH</Text>
    </Flex>
  )
}

export default WalletDetail
