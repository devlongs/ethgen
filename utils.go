package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func printText(account Account) {
	fmt.Println("Private Key:", account.PrivateKey)
	fmt.Println("Public Key:", account.PublicKey)
	fmt.Println("Ethereum Address:", account.Address)
}

func printJSON(account Account) {
	output, err := json.MarshalIndent(account, "", "  ")
	if err != nil {
		log.Fatal("Error formatting JSON:", err)
	}
	fmt.Println(string(output))
}
