package main

import (
	"net/http"
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
