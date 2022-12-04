package main

import (
	"fmt"
	"net/http"
	"net/url"
)

/**
 * Directors is a list of functions that modify the
 * request before it is sent to the target server.
 **/
var Directors []func(req *http.Request)

/**
  * ModifyResponses is a list of functions that modify the
  * response before it is sent to the client.
**/
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
		url, _ := url.Parse(TARGET_SERVICE_BASE_URL + req.URL.Path)
		req.URL = url
		fmt.Printf("proxy to %v\n", url)
	}

	Directors = []func(req *http.Request){simpleDirector}
}
