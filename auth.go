package main

import (
	"net/http"
	"regexp"
)

var re = regexp.MustCompile(".+\\.(html|css|js|jpg|png|gif)")

func passIfCleanContent(req *http.Request) bool {
	if req.Method != "GET" {
		return false
	}

	if req.RequestURI == "/" || re.MatchString(req.URL.Path) {
		req.URL.Path = re.FindString(req.URL.Path)
		return true
	}

	return false
}

func authenticateRequest(req *http.Request) bool {
	// todo: valid cookie
	rawCookie := req.Header.Get("Cookie")
	cookies := parseCookies(rawCookie)
	tokenCookie, err := getCookie("bullet-token", cookies)

	if err != nil {
		return false
	}

	return validateToken(tokenCookie.Value)
}

func validateToken(_ string) bool {
	return true
}
