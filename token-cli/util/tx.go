package util

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

type TxData struct {
	Hash  string
	Block *big.Int
}

// this function wraps mutating transactions and allows to control their flow,
// for example wait for block mining or just send request and give you potential hash
//
// waitMined - flag which allows to wait for block with tx to be mined
// action - is actual smart contract action you want to execute
func WrapTx(waitMined bool, ctx context.Context, conn bind.DeployBackend, action func() (*types.Transaction, error)) (*TxData, error) {
	tx, err := action()
	if err != nil {
		return nil, fmt.Errorf("failed to perform action in tx wrapper: %w", err)
	}

	// return calculated hash if dont wait
	if !waitMined {
		return &TxData{
			Hash: tx.Hash().String(),
		}, nil
	}

	rec, err := bind.WaitMined(ctx, conn, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to wait mined tx: %w", err)
	}

	return &TxData{
		Hash:  rec.BlockHash.String(),
		Block: rec.BlockNumber,
	}, nil

}
