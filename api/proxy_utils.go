package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

/**
 * LogReq is the function that logs the request.
**/
func LogReq(req *http.Request) {
	fmt.Printf("[%s] %s%s\n=> %s%s\n\n", req.Method, req.Host, req.RequestURI, req.URL.Host, req.URL.Path)
}

/**
 * GoToErrorPage is the function that redirects to the error page.
 */
func GoToErrorPage(msg string, err error, req *http.Request) {
	fmt.Fprintln(os.Stderr, msg, err)

	errorPageURL, err := url.Parse(ERROR_PAGE_URL)

	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}

	req.URL.Scheme = errorPageURL.Scheme
	req.URL.Host = errorPageURL.Host
	req.URL.Path = "/404?msg=" + msg

	LogReq(req)
}
