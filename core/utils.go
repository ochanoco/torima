package main

import (
	"net/http"
)

func delCookie(name string, src []*http.Cookie) []*http.Cookie {
	var dest []*http.Cookie

	for _, v := range src {
		if v.Name != name {
			dest = append(dest, v)
		}
	}

	return dest
}
