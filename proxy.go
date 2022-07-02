package main

import (
	"fmt"
	"net/http"
	"net/url"
)

var Directors []func(req *http.Request)
var ModifyResponses []func(req *http.Response)

func director(req *http.Request) {
	for _, d := range Directors {
		d(req)
	}
}

func modifyResponse(res *http.Response) error {
	for _, mR := range ModifyResponses {
		mR(res)
	}

	return nil
}

func init() {
	// Client sets the original URL in the Ochanoco-Forward-For header
	simpleDirector := func(req *http.Request) {
		url, _ := url.Parse(SERVWICE_BASE_URL)
		req.URL = url
		fmt.Printf("proxy to %v\n", url)
	}

	Directors = []func(req *http.Request){simpleDirector}
}
