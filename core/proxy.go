package main

import (
	"net/http"
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
