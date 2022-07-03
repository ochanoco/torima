package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEncAndDec(t *testing.T) {
	t.Run("enc/dec", func(t *testing.T) {
		iv := []byte("aaaaaaaaaaaaaaaa")
		key := []byte("bbbbbbbbbbbbbbbb")
		plaintext := []byte("plaintext")

		encToken := enc(plaintext, key, iv)
		decToken := dec(encToken, key, iv)

		if bytes.Compare(plaintext, decToken) != 0 {
			msg := fmt.Sprintf("wrong response: '%s'\nexpected: '%s'", decToken, plaintext)
			t.Error(msg)
		}
	})
}
