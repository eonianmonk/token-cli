# token-cli
token smart contract and  golang cli app to control it

# Requirements: 
- node version >= 20 (for hardhat)
- go version >= 1.22
- abigen >= 1.10
- installed docker

# Smart contracts
Smart contracts are located in solidity/contracts. There are imported from openzeppelin Ownalbe, locally implemanted abstract contract Freezable which stores data about frozen accounts, interface ITestToken which describes token related functionality and finally implementation of TestToken in TestToken.sol.

Tests are conducted using [hardhat](https://hardhat.org/docs). 

To deploy Smart contract create .env and fill it with your data (your private key and rpc endpoint).