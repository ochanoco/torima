package main

import (
	"net/http"
)

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
