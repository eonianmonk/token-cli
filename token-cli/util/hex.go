package util

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/alecthomas/kong"
	"github.com/ethereum/go-ethereum/crypto"
)

// eth address
type HexBytes []byte

// decoder for kong
func (h *HexBytes) Decode(ctx *kong.DecodeContext) error {

	token, err := ctx.Scan.PopValue("")
	if err != nil {
		return fmt.Errorf("failed to retrieve string value for hexbyte: %w", err)
	}
	value, convOk := token.Value.(string)
	if !convOk {
		return fmt.Errorf("failed to convert token value to string")
	}
	// zero address, for burning tx
	if value == "0" {
		*h = make([]byte, 20)
		return nil
	}

	decoded, err := HexToBytes(value)
	if err != nil {
		return fmt.Errorf("failed to parse HexBytes: %w", err)
	}

	if len(decoded) != 20 {
		return fmt.Errorf("invalid address len: got %d, expected 20", len(decoded))
	}

	*h = decoded
	return nil
}

// decoder for yaml
func (h *HexBytes) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var data string
	err := unmarshal(&data)
	if err != nil {
		return err
	}

	decoded, err := HexToBytes(data)
	if err != nil {
		return err
	}

	*h = decoded

	return nil
}

func (h HexBytes) String() string {
	return hex.EncodeToString(h)
}

// converts hexadecimal string to bytes
func HexToBytes(hexadecimal string) ([]byte, error) {
	hexadecimal = strings.TrimPrefix(hexadecimal, "0x")
	return hex.DecodeString(hexadecimal)
}

func PrivateKeyFromHex(hexadecimal string) (*ecdsa.PrivateKey, error) {
	hexadecimal = strings.TrimPrefix(hexadecimal, "0x")
	pkBytes, err := HexToBytes(hexadecimal)
	if err != nil {
		return nil, fmt.Errorf("failed to decode private key from env: %w", err)
	}

	return crypto.ToECDSA(pkBytes)

}
