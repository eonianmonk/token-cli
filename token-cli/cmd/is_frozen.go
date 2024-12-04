package main

import (
	"context"
	"fmt"
	testtoken "token-cli/services/smart_contracts/test_token"
	"token-cli/util"

	"github.com/ethereum/go-ethereum/common"
)

type IsFrozen struct {
	Address util.HexBytes `required:"" help:"account address"`
}

func (cmd *IsFrozen) Run(cfg *GlobalCfg) error {
	callOpts := testtoken.NewCallOpts(context.Background(), common.Address{})

	account := common.Address(cmd.Address)

	isFrozen, err := cfg.Contract.IsFrozen(callOpts, account)
	if err != nil {
		fmt.Printf("failed to get balance: %e", err)
	}

	fmt.Printf("%s frozen state is %t", account.Hex(), isFrozen)

	return nil
}
