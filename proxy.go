package main

import (
	"net/http"
)

var directors []func(req *http.Request)
var modifyResponses []func(req *http.Response)

func director(req *http.Request) {
	for _, d := range directors {
		d(req)
	}
}

func modifyResponse(res *http.Response) error {
	for _, mR := range modifyResponses {
		mR(res)
	}

	return nil
}
