package certs

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"log"
)

var (
	privateKey *rsa.PrivateKey
)

func init() {
	var err error
	privateKey, err = rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		e := fmt.Errorf("init: failed to generate the key pair: %w", err)
		panic(e)
	}

	buff := EncodePubkeyPem(&privateKey.PublicKey)
	log.Printf("init: generated public key:\n%v", string(buff))
}

func Sign(data []byte) ([]byte, error) {
	hashed := sha256.Sum256(data)

	sig, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		e := fmt.Errorf("Sign: failed to sign data: %w", err)
		return nil, e
	}

	return sig, nil
}

func Verify(msg []byte, signature []byte) bool {
	hashed := sha256.Sum256(msg)

	err := rsa.VerifyPKCS1v15(&privateKey.PublicKey, crypto.SHA256, hashed[:], signature)
	if err != nil {
		log.Printf("Verify: invalid signature")
		return false
	}

	return true
}
