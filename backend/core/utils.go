package core

import (
	"crypto/rand"
	"encoding/json"
	"net/http"
)

func RandomString(length int) (string, error) {
	const chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	bytes := make([]byte, length)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	for i, b := range bytes {
		bytes[i] = chars[b%byte(len(chars))]
	}

	return string(bytes), nil
}

func DumpHeader(headers http.Header) (string, error) {
	b, err := json.Marshal(headers)
	s := string(b)
	return s, err
}
