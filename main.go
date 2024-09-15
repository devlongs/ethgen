package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"
)

type Account struct {
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
	Address    string `json:"address"`
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "ethgen",
		Short: "Ethereum account generator",
		Run:   generateAccount,
	}

	rootCmd.Flags().StringP("format", "f", "text", "Output format: text or json")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func generateAccount(cmd *cobra.Command, args []string) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal("Error generating private key:", err)
	}

	publicKey := privateKey.PublicKey

	publicAddress := crypto.PubkeyToAddress(publicKey)

	account := Account{
		PrivateKey: privateKeyHex(privateKey),
		PublicKey:  publicKeyHex(&publicKey),
		Address:    publicAddress.Hex(),
	}

	format, _ := cmd.Flags().GetString("format")
	switch format {
	case "json":
		printJSON(account)
	default:
		printText(account)
	}
}

func privateKeyHex(privateKey *ecdsa.PrivateKey) string {
	return fmt.Sprintf("0x%x", crypto.FromECDSA(privateKey))
}

func publicKeyHex(publicKey *ecdsa.PublicKey) string {
	return fmt.Sprintf("0x%x", crypto.FromECDSAPub(publicKey))
}
