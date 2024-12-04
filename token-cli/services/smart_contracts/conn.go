package smartcontracts

import (
	"fmt"
	token_cli "token-cli"

	"github.com/ethereum/go-ethereum/ethclient"
)

func NewClient(netcfg token_cli.NetworkConfig) (*ethclient.Client, error) {
	conn, err := ethclient.Dial(netcfg.RPC)
	if err != nil {
		return nil, fmt.Errorf("failed to dial rpc: %w", err)
	}
	return conn, nil
}
