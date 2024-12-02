# token-cli
token smart contract and  golang cli app to control it

# Quickstart
Requirements: 
- node version >= 20 (for hardhat)

# Smart contracts
Smart contracts are located in solidity/contracts. There are imported from openzeppelin Ownalbe, locally implamanted abstract contract Freezable which stores data about frozen accounts, interface ITestToken which describes token related functionality and finally implementation of TestToken in TestToken.sol.
Tests are conducted using [hardhat](https://hardhat.org/docs). 