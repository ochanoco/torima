package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)


/**
 * GoToErrorPage is the function that redirects to the error page.
 */
func GoToErrorPage(msg string, err error, req *http.Request) {
	fmt.Fprintln(os.Stderr, msg, err)

	errorPageURL, err := url.Parse(ERROR_URL)

	if err != nil {
		log.Fatalf("%s: %v\n\n", msg, err)
	}

	req.URL.Scheme = errorPageURL.Scheme
	req.URL.Host = errorPageURL.Host
	req.URL.Path = "/404?msg=" + msg

	LogReq(req)
}
