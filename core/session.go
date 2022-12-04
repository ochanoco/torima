package main

import (
	"net/http"
)

/**
 * RemoveToken is a director that removes the token
**/
func RemoveToken(req *http.Request) {
	queries := req.URL.Query()
	queries.Del("token")

	req.URL.RawQuery = queries.Encode()

	cookies := req.Cookies()
	cookies = delCookie("token", cookies)

	for _, cookie := range cookies {
		req.Header.Set("Set-Cookie", cookie.String())
	}
}
