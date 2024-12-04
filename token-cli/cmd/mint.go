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

type Mint struct {
	Receiver util.HexBytes `name:"to" required:"" help:"mint token receiver address"`
	Amount   int           `required:"" help:"amount to mint"`
}

func (cmd *Mint) Run(cfg *GlobalCfg) error {
	txOpts, err := testtoken.NewTxOpts(cfg.Network.ChainId, cfg.PrivateKey)
	if err != nil {
		return err
	}

	receiver := common.Address(cmd.Receiver)
	amount := big.NewInt(int64(cmd.Amount))

	res, err := util.WrapTx(cfg.Globals.WaitMined, context.Background(), cfg.Client,
		func() (*types.Transaction, error) {
			return cfg.Contract.Mint(txOpts, receiver, amount)
		})
	if err != nil {
		return err
	}

	fmt.Printf("minted %s to %s. Tx: %s\n", amount.String(), receiver.String(), res.Hash)
	if res.Block != nil {
		fmt.Printf("Block: %s", res.Block.String())
	}

	return nil
}
