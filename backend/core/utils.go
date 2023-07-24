package core

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
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
