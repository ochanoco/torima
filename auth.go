package main

import "net/http"

func authenticateRequest(req *http.Request) bool {
	if req.Method == "GET" && req.RequestURI == "/" {
		return true
	}

	// todo: valid cookie
	rawCookie := req.Header.Get("Cookie")
	cookies := parseCookies(rawCookie)
	tokenCookie, err := getCookie("tauset-token", cookies)

	if err != nil {
		return false
	}

	return validateToken(tokenCookie.Value)
}

func validateToken(_ string) bool {
	return true
}
