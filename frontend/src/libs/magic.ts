import { Magic } from "magic-sdk"

// Initialize the Magic instance
export const magic = new Magic("pk_live_F4A72E6ABE77B965", {
  network: {
    rpcUrl: process.env.NETWORK_RPC_URL || "https://rpc-mumbai.matic.today", // TODO: .env config not working!?!?
    chainId: 80001, // Mumbai
  },
})
magic.network = 'matic';
