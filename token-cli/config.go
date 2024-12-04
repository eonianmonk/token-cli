package token_cli

import (
	"fmt"
	"os"
	"token-cli/util"

	"gopkg.in/yaml.v2"
)

const (
	ConfigFileName         = "config.yaml"
	SenderPrivateKeyEnvKey = "DEPLOYMENT_PRIVATE_KEY"
)

func LoadConfig() (*Config, error) {

	f, err := os.Open(ConfigFileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file %s: %w", ConfigFileName, err)
	}

	d := yaml.NewDecoder(f)
	cfg := &Config{}

	err = d.Decode(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to decode yaml config: %w", err)
	}

	cfg.SenderPK = os.Getenv(SenderPrivateKeyEnvKey)

	return cfg, nil
}

type Config struct {
	Networks map[string]NetworkConfig `yaml:"networks"`
	SenderPK string                   `env:"DEPLOYMENT_PRIVATE_KEY"`
}

type NetworkConfig struct {
	RPC             string        `yaml:"rpc"`
	ChainId         int           `yaml:"chain_id"`
	ContractAddress util.HexBytes `yaml:"contract_address"`
}
