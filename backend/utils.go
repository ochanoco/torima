package main

import (
	"crypto/rand"
	"net/http"
)

/**
 * delCookie is a utils function that remove the user token.
 **/

func delCookie(name string, src []*http.Cookie) []*http.Cookie {
	var dest []*http.Cookie

	for _, v := range src {
		if v.Name != name {
			dest = append(dest, v)
		}
	}

	return dest
}

func randomString(length int) (string, error) {
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
