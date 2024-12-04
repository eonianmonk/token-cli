# token-cli
token smart contract and  golang cli app to control it

# Requirements: 
- node version >= 20 (for hardhat)
- go version >= 1.22
- abigen >= 1.10
- installed docker

# Quickstart 


1. Install NPM dependencies  

```make install``` 

2. Build ABI and generate Go code from ABI

```make build```

3. Build docker image for cli

```make build-docker```

4. Change ./cli.sh to executable 

```chmod +x ./cli.sh```

5. To start interacting with smart contract create .env file and fill it according to .env.example (private keys in hexadecimal; also you may omit holesky rps endpoint and deployment private key if you dont plan to deploy smart contract, sender_private_key is sufficient for running cli app).

If you want to interact with your deployed smart contract or on another chain you may modify your access in config.yaml. Also this config supports multiple networks simultaneously and you may specify the target network/contract with --network flag for cli app.

It must be enough and now you can use ./cli.sh, for available commands you can run ```./cli.sh --help``` or ```sudo ./cli.sh --help``` if your docker does not work without sudo

To run smart contract tests you can use ```make test```

To deploy smart contract yourself use ```make deploy-token```



# Smart contract
Predeployed smart contract is at ```0x41B471B930CcdBe2Ea2D4114Bb95679d8B781e7b```.

Smart contracts are located in solidity/contracts. 

Tests are conducted using [hardhat](https://hardhat.org/docs). 

To deploy Smart contract create .env and fill it with your data (your private key and rpc endpoint).

# CLI app

I used [kong](https://github.com/alecthomas/kong) for CLI and [go-ethereum](https://github.com/ethereum/go-ethereum) for smart contract interaction.
 
All mutating commands output transaction hash. If you use --wait-mined flag, you'll also get block number.

And using flag --privateKey you may override private key from .env.


# Potential improvements 

1. There should be only one config, not some parts in .env and some in config.yaml. 


### CLI help 

```
me@mypc:~$ sudo ./cli.sh --help
Usage: token-cli <command> [flags]

Flags:
  -h, --help                 Show context-sensitive help.
      --privateKey=STRING    sender's private key overwriting env private key
  -n, --network="holesky"    omit to use default (holesky)
      --wait-mined           wait for tx mined and also get a block number

Commands:
  isFrozen --address=HEX-BYTES [flags]
    check if account is frozen

  freeze --address=HEX-BYTES [flags]
    freeze account

  unfreeze --address=HEX-BYTES [flags]
    unfreeze account

  mint --to=HEX-BYTES --amount=INT [flags]
    mint to account

  balanceOf --address=HEX-BYTES [flags]
    sget account balance

  transfer --to=HEX-BYTES --amount=INT [flags]
    transfer tokens

Run "token-cli <command> --help" for more information on a command.
```

# Usage examples

Minting to `0x4770B19C113cbC07e220d51E6A9fBC4ED30Fa51a`:

```
me@mypc:~$ ./cli.sh mint --to 0x4770B19C113cbC07e220d51E6A9fBC4ED30Fa51a --amount 1111
Using 0xea8561756516A3f48d3eCF2229A8a07F5916a39c private key
minted 1111 to 0x4770B19C113cbC07e220d51E6A9fBC4ED30Fa51a. Tx: 0xc8d962602bf2dd741a5588f09432989430aa50867eb619adf812beab8b951a03

me@mypc:~$ ./cli.sh balanceOf --address 0x4770B19C113cbC07e220d51E6A9fBC4ED30Fa51a
Using 0xea8561756516A3f48d3eCF2229A8a07F5916a39c private key
0x4770B19C113cbC07e220d51E6A9fBC4ED30Fa51a balance is 1111
```

### freeze/unfreeze:

```
me@mypc:~$ ./cli.sh freeze --address 0xea8561756516A3f48d3eCF2229A8a07F5916a39c
Using 0xea8561756516A3f48d3eCF2229A8a07F5916a39c private key
frozen 0xea8561756516A3f48d3eCF2229A8a07F5916a39c. Tx: 0xf93a6cb3b0f45f63889a4b632d7714da6f52ebb1fb3aa72bf30072e70a179c43

me@mypc:~$ ./cli.sh transfer --to 0xea8561756516A3f48d3eCF2229A8a07F5916a39c --amount 1 --privateKey f55ee2bc4df60073ba3376bf1d1a38aac5e7218a268c45fafaeedfb012810637
Using 0xfe2D9AFe510Cc1BB07e5D4f62b85f737CE1cCA6A private key
failed to do transfer: failed to perform action in tx wrapper: execution reverted: account is frozen

me@mypc:~$ ./cli.sh unfreeze --address 0xea8561756516A3f48d3eCF2229A8a07F5916a39c
Using 0xea8561756516A3f48d3eCF2229A8a07F5916a39c private key
unfrozen 0xea8561756516A3f48d3eCF2229A8a07F5916a39c. Tx: 0x7e66816285cee6e6cbd161f7f5955f2df50d0e51d37ea212f85373051a314c81

me@mypc:~$ ./cli.sh transfer --to 0xfe2D9AFe510Cc1BB07e5D4f62b85f737CE1cCA6A --amount 1
Using 0xea8561756516A3f48d3eCF2229A8a07F5916a39c private key
transfered 1 to 0xfe2D9AFe510Cc1BB07e5D4f62b85f737CE1cCA6A. Tx: 0x5d316700a0bb635c8ee40a3c42b633fd675f4a1b1469652e667596d7e57b5e9c
```

### --wait-mined and transfers

```command
me@mypc:~$ ./cli.sh balanceOf --address 0xfe2D9AFe510Cc1BB07e5D4f62b85f737CE1cCA6A
Using 0xea8561756516A3f48d3eCF2229A8a07F5916a39c private key
0xfe2D9AFe510Cc1BB07e5D4f62b85f737CE1cCA6A balance is 8

me@mypc:~$ ./cli.sh balanceOf --address 0xea8561756516A3f48d3eCF2229A8a07F5916a39c
Using 0xea8561756516A3f48d3eCF2229A8a07F5916a39c private key
0xea8561756516A3f48d3eCF2229A8a07F5916a39c balance is 97

me@mypc:~$ ./cli.sh transfer --to 0xfe2D9AFe510Cc1BB07e5D4f62b85f737CE1cCA6A --amount 1
Using 0xea8561756516A3f48d3eCF2229A8a07F5916a39c private key
transfered 1 to 0xfe2D9AFe510Cc1BB07e5D4f62b85f737CE1cCA6A. Tx: 0xb1bceab00e89a1d5a520376cae42e8e76c9a4a115c89d0e8b4ef5e2da1125af9

me@mypc:~$ ./cli.sh balanceOf --address 0xea8561756516A3f48d3eCF2229A8a07F5916a39c
Using 0xea8561756516A3f48d3eCF2229A8a07F5916a39c private key
0xea8561756516A3f48d3eCF2229A8a07F5916a39c balance is 96

me@mypc:~$ ./cli.sh transfer --to 0xfe2D9AFe510Cc1BB07e5D4f62b85f737CE1cCA6A --amount 1 --wait-mined
Using 0xea8561756516A3f48d3eCF2229A8a07F5916a39c private key
transfered 1 to 0xfe2D9AFe510Cc1BB07e5D4f62b85f737CE1cCA6A. Tx: 0x9d81b4310f0d6005d646c8aaec1a9d5e148a5025207998527f4c4012122de6dd
Block: 2868323

me@mypc:~$ ./cli.sh balanceOf --address 0xea8561756516A3f48d3eCF2229A8a07F5916a39c
Using 0xea8561756516A3f48d3eCF2229A8a07F5916a39c private key
0xea8561756516A3f48d3eCF2229A8a07F5916a39c balance is 95

me@mypc:~$ ./cli.sh balanceOf --address 0xfe2D9AFe510Cc1BB07e5D4f62b85f737CE1cCA6A
Using 0xea8561756516A3f48d3eCF2229A8a07F5916a39c private key
0xfe2D9AFe510Cc1BB07e5D4f62b85f737CE1cCA6A balance is 10
```

### --privateKey

```command
me@mypc:~$ ./cli.sh transfer --to 0xea8561756516A3f48d3eCF2229A8a07F5916a39c --amount 1 --privateKey (0xfe's private key)
Using 0xfe2D9AFe510Cc1BB07e5D4f62b85f737CE1cCA6A private key
transfered 1 to 0xea8561756516A3f48d3eCF2229A8a07F5916a39c. Tx: 0xa5139f623804d96380761adc760d0ce8f3061c145320ab387867be574a362cd7

me@mypc:~$ ./cli.sh balanceOf --address 0xea8561756516A3f48d3eCF2229A8a07F5916a39c
Using 0xea8561756516A3f48d3eCF2229A8a07F5916a39c private key
0xea8561756516A3f48d3eCF2229A8a07F5916a39c balance is 96

me@mypc:~$ ./cli.sh balanceOf --address 0xfe2D9AFe510Cc1BB07e5D4f62b85f737CE1cCA6A
Using 0xea8561756516A3f48d3eCF2229A8a07F5916a39c private key
0xfe2D9AFe510Cc1BB07e5D4f62b85f737CE1cCA6A balance is 9
```