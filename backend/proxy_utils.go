package main

import (
	"fmt"
	"net/http"
	"os"
)

/**
 * GoToErrorPage is the function that redirects to the error page.
 */
func GoToErrorPage(msg string, err error, req *http.Request) {
	fmt.Fprintln(os.Stderr, msg, err)

	req.URL.Scheme = ErrorUrl.Scheme
	req.URL.Host = ErrorUrl.Host
	req.URL.Path = "/404?msg=" + msg

	LogReq(req)
}
