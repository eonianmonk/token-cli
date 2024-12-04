package main

import (
	"crypto/ecdsa"
	token_cli "token-cli"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type GlobalCfg struct {
	Globals    Globals
	Network    token_cli.NetworkConfig
	Contract   *token_cli.TestToken
	PrivateKey *ecdsa.PrivateKey

	// required to control flow
	Client bind.DeployBackend
}

type Globals struct {
	PrivateKey string `name:"privateKey" optional:"" help:"sender's private key overwriting env private key"`
	Network    string `short:"n" default:"holesky" help:"omit to use default (holesky)"`
	WaitMined  bool   `help:"wait for tx mined and also get a block number"`
}

type CLI struct {
	Globals

	IsFrozen  IsFrozen  `cmd:"" name:"isFrozen" help:"check if account is frozen"`
	Freeze    Freeze    `cmd:"" name:"freeze" help:"freeze account"`
	Unfreeze  Unfreeze  `cmd:"" name:"unfreeze" help:"unfreeze account"`
	Mint      Mint      `cmd:"" name:"mint" help:"mint to account"`
	BalanceOf BalanceOf `cmd:"" name:"balanceOf" help:"sget account balance"`
	Transfer  Transfer  `cmd:"" name:"transfer" help:"transfer tokens"`
}
