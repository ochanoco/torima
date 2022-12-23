package main

import (
	"net/http"
	"regexp"
)

var re = regexp.MustCompile(".+\\.(html|css|js|jpg|png|gif)")

/**
 * passIfCleanContent is function that is it checked authenticated needed.
 * If the request is not GET or the request is not for static content, the request is passed.
**/

func passIfCleanContent(req *http.Request) bool {
	if req.Method != "GET" {
		// If the request is not GET, the request is passed.
		return false
	}

	if req.RequestURI == "/" || re.MatchString(req.URL.Path) {
		// If the request is for static content, the request is passed.
		req.URL.Path = re.FindString(req.URL.Path)
		return true
	}

	return false
}

/**
 * authenticateRequest is a function for user authentication.
 * Validate the token on cookie and if it is valid, pass the request.
 */
func authenticateRequest(req *http.Request) bool {
	cookie, err := req.Cookie("ochanoco-token")

	if err != nil {
		return false
	}

	return validateToken(cookie.Value)
}

/**
 * TODO: check token
 */
func validateToken(_ string) bool {
	return true
}
