package main

import (
	"os"

	"github.com/safe2008/vault-actions/pkg/vault"
)

func main() {
	address := os.Getenv("INPUT_ADDRESS")
	token := os.Getenv("INPUT_TOKEN")
	secrets := os.Getenv("INPUT_SECRETS")
	vault.GetSecret(address, token, secrets)
}
