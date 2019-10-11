package main

import (
	"fmt"

	"github.com/inn4science/exonum-go"
	"github.com/inn4science/exonum-go/crypto"
	"github.com/inn4science/exonum-go/internal/testTypes"
)

func main() {
	kp := crypto.KeyPair{}.Random()
	sk := kp.SecretKey()
	pk := kp.PublicKey()
	fmt.Println("SecretKey:", sk.String())
	fmt.Println("PublicKey:", pk.String())

	serviceID := uint16(0)
	messageID := uint16(1)

	tx := exonum.ServiceTx{}.New(&testTypes.Transfer{
		From:   "alice",
		To:     "bob",
		Amount: 100,
		Seed:   42,
	}, kp.PublicKey(), serviceID, messageID)

	sig, err := tx.IntoSignedTx(*kp.SecretKey())
	fmt.Println("sig err:", err)
	fmt.Println("SignedTx:", sig)

	hash, err := tx.Hash()
	fmt.Println("hash err:", err)
	fmt.Println("TxHash:", hash)
}
