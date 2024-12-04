package smartcontracts

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	token_cli "token-cli"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func NewTxOpts(networkCfg token_cli.NetworkConfig, privateKey *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
	chainId := big.NewInt(int64(networkCfg.ChainId))

	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return nil, fmt.Errorf("failed to create keyed transactor: %w", err)
	}

	return opts, nil
}
