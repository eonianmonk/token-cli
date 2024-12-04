package main

import (
	"context"
	"fmt"
	testtoken "token-cli/services/smart_contracts/test_token"
	"token-cli/util"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Freeze struct {
	Address util.HexBytes `required:"" help:"account address"`
}

func (cmd *Freeze) Run(cfg *GlobalCfg) error {
	txOpts, err := testtoken.NewTxOpts(cfg.Network.ChainId, cfg.PrivateKey)
	if err != nil {
		return err
	}

	account := common.Address(cmd.Address)

	res, err := util.WrapTx(cfg.Globals.WaitMined, context.Background(), cfg.Client,
		func() (*types.Transaction, error) {
			return cfg.Contract.Freeze(txOpts, account)
		})
	if err != nil {
		return err
	}

	fmt.Printf("frozen %s. Tx: %s\n", account.String(), res.Hash)
	if res.Block != nil {
		fmt.Printf("Block: %s", res.Block.String())
	}

	return nil
}
