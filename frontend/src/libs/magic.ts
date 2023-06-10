import { Magic } from "magic-sdk"

// Initialize the Magic instance
export const magic = new Magic("pk_live_F4A72E6ABE77B965", {
  network: {
    rpcUrl: "https://eth-sepolia.g.alchemy.com/v2/demo",
    chainId: 11155111,
  },
})
