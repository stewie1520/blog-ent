//go:build paseto_gen_key

package main

import (
	"fmt"

	"aidanwoods.dev/go-paseto"
)

func main() {
	key := paseto.NewV4AsymmetricSecretKey()
	fmt.Printf("Private key: %s\n", key.ExportHex())
	fmt.Printf("Public key: %s\n", key.Public().ExportHex())
}
