package core

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ochanoco/proxy/ent"
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

func ReadHTTPBody(body *io.ReadCloser) ([]byte, error) {
	bodyBuf, err := io.ReadAll(*body)
	if err != nil {
		return nil, fmt.Errorf("RequestLogDirector: non-nil error while reading request body: %v", err)
	}

	(*body).Close()
	b := io.NopCloser(bytes.NewBuffer(bodyBuf))
	*body = b

	return bodyBuf, err
}

type HashChain = ent.HashChain

func CalcHashChain(log *ent.ServiceLog, previous *ent.HashChain) []byte {
	var previous_buf []byte

	if previous_buf == nil {
		previous_buf = []byte("invalid")
	} else {
		previous_buf = previous.Hash
	}

	hasher := crypto.SHA256.New()
	hasher.Write([]byte(log.Headers))
	hasher.Write(log.Body)
	hasher.Write(previous_buf)
	hash := hasher.Sum(nil)

	return hash
}
