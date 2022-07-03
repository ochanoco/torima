package main

import (
	"crypto/aes"
	"crypto/cipher"
	"log"
)

func enc(plaintext []byte, key []byte, iv []byte) []byte {
	c, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("Error: NewCipher(%d bytes) = %s", len(key), err)
	}

	cfb := cipher.NewCFBEncrypter(c, iv)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)

	return ciphertext
}

func dec(plaintext []byte, key []byte, iv []byte) []byte {
	c, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("Error: NewCipher(%d bytes) = %s", len(key), err)
	}

	cfb := cipher.NewCFBDecrypter(c, iv)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)

	return ciphertext
}
