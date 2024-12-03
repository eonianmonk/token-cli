/** @type import('hardhat/config').HardhatUserConfig */
require("@nomicfoundation/hardhat-toolbox");
require('ethers')
require('dotenv').config({path:'../.env'})

const DEFAULT_HOLESKY_RPC = "https://ethereum-holesky.publicnode.com"

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.27",
  paths:{
    sources:"./contracts",
    tests:"./tests",
  },
  // https://hardhat.org/hardhat-runner/docs/config#networks-configuration
  networks: {
    holesky: {
      url: process.env.HOLESKY_RPC_ENDPOINT || DEFAULT_HOLESKY_RPC,
      accounts: process.env.DEPLOYMENT_PRIVATE_KEY? [process.env.DEPLOYMENT_PRIVATE_KEY]:[],
    }
  }
};

// console.log( module.exports.networks.holesky.accounts )