package main

import (
	"context"
	"fmt"
	testtoken "token-cli/services/smart_contracts/test_token"
	"token-cli/util"

	"github.com/ethereum/go-ethereum/common"
)

type BalanceOf struct {
	Address util.HexBytes `required:"" help:"account address"`
}

func (cmd *BalanceOf) Run(cfg *GlobalCfg) error {
	callOpts := testtoken.NewCallOpts(context.Background(), common.Address{})

	account := common.Address(cmd.Address)

	balance, err := cfg.Contract.BalanceOf(callOpts, account)
	if err != nil {
		fmt.Printf("failed to get balance: %e", err)
	}

	fmt.Printf("%s balance is %s", account.Hex(), balance.String())

	return nil
}
