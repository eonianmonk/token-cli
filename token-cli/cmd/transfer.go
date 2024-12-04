package main

import (
	"context"
	"fmt"
	"math/big"
	testtoken "token-cli/services/smart_contracts/test_token"
	"token-cli/util"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Transfer struct {
	Receiver util.HexBytes `required:"" name:"to" help:"receiver's address"`
	Amount   int           `required:"" help:"amount to transfer. must be >= balance"`
}

func (cmd *Transfer) Run(cfg *GlobalCfg) error {
	txOpts, err := testtoken.NewTxOpts(cfg.Network.ChainId, cfg.PrivateKey)
	if err != nil {
		return err
	}

	receiver := common.Address(cmd.Receiver)
	amount := big.NewInt(int64(cmd.Amount))

	res, err := util.WrapTx(cfg.Globals.WaitMined, context.Background(), cfg.Client,
		func() (*types.Transaction, error) {
			return cfg.Contract.Transfer(txOpts, receiver, amount)
		})
	if err != nil {
		return err
	}

	fmt.Printf("transfered %s to %s. Tx: %s\n", amount.String(), receiver.String(), res.Hash)
	if res.Block != nil {
		fmt.Printf("Block: %s", res.Block.String())
	}

	return nil
}
