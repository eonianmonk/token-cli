package testtoken

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	token_cli "token-cli"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func NewContract(networkCfg token_cli.NetworkConfig, conn *ethclient.Client) (*token_cli.TestToken, error) {

	contractAddress := common.Address(networkCfg.ContractAddress)

	txer, err := token_cli.NewTestToken(contractAddress, conn)
	if err != nil {
		return nil, err
	}

	return txer, nil
}

func NewTxOpts(chainId int, pk *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
	chainIdBigint := big.NewInt(int64(chainId))

	opts, err := bind.NewKeyedTransactorWithChainID(pk, chainIdBigint)
	if err != nil {
		return nil, fmt.Errorf("failed to create opts: %w", err)
	}

	return opts, nil
}

func NewCallOpts(ctx context.Context, from common.Address) *bind.CallOpts {
	return &bind.CallOpts{Context: ctx, From: from}
}
