package main

import (
	"fmt"
	"os"
	token_cli "token-cli"
	smartcontracts "token-cli/services/smart_contracts"
	testtoken "token-cli/services/smart_contracts/test_token"
	"token-cli/util"

	"github.com/alecthomas/kong"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	var cli CLI
	ctx := kong.Parse(&cli,
		kong.ShortUsageOnError(),
	)

	cfg, err := token_cli.LoadConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	network, ok := cfg.Networks[cli.Network]
	if !ok {
		fmt.Printf("undefined network %s", cli.Network)
		os.Exit(1)
	}

	conn, err := smartcontracts.NewClient(network)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	contract, err := testtoken.NewContract(network, conn)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	// cli private key overrides config pk
	if len(cli.PrivateKey) != 0 {
		cfg.SenderPK = cli.PrivateKey
	}

	pk, err := util.PrivateKeyFromHex(cfg.SenderPK)
	if err != nil {
		fmt.Printf("failed to convert sender's PK to crypto.ECDSA PK: %s", err)
		os.Exit(1)
	}

	fmt.Printf("Using %s private key\n", crypto.PubkeyToAddress(pk.PublicKey))

	err = ctx.Run(&GlobalCfg{
		Globals:    cli.Globals,
		Network:    network,
		Contract:   contract,
		PrivateKey: pk,
		Client:     conn,
	})
	if err != nil {
		fmt.Printf("failed to do %s: %s", ctx.Command(), err)
		os.Exit(1)
	}

	os.Exit(0)
}
