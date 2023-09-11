import "@openzeppelin/hardhat-upgrades";
import { HardhatUserConfig, task, types } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import { config } from "dotenv";

config();

const hhconfig: HardhatUserConfig = {
  solidity: {
    version: "0.8.19",
    settings: {
      evmVersion: "paris",
      optimizer: {
        enabled: true,
        runs: 200,
      },
    },
  },
  paths: {
    sources: "./src"
  },
  networks: {
    polygon: {
      url: "https://polygon-rpc.com/",
      accounts: ["4bc869b662589948210bba09d83fcc3fa809c19afbad55e6bb45d7745bbf01c3" || "0x0"],
    },
  },
  etherscan: {
    apiKey: {
      polygon: "XB9IQXXNRQJE8AAQHNW4TWFZX2JBJ3DWJE" || ""
    },
  },
};

export default hhconfig;
