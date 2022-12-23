package main

import (
	"net/http"
	"net/http/httputil"
)

type OchanocoDirector = func(proxy *OchanocoProxy, req *http.Request)
type OchanocoResponse = func(proxy *OchanocoProxy, req *http.Response)

type OchanocoProxy struct {
	Directors       []OchanocoDirector
	ModifyResponses []OchanocoResponse
	ReverseProxy    *httputil.ReverseProxy
	Database        *Database
}

/**
 * Directors is a list of functions that modify the
 * request before it is sent to the target server.
 **/
func (proxy *OchanocoProxy) Director(req *http.Request) {
	for _, d := range proxy.Directors {
		d(proxy, req)
	}
}

/**
  * ModifyResponses is a list of functions that modify the
  * response before it is sent to the client.
**/
func (proxy *OchanocoProxy) ModifyResponse(res *http.Response) error {
	for _, mR := range proxy.ModifyResponses {
		mR(proxy, res)
	}

	return nil
}

func (proxy *OchanocoProxy) AddDirector(director OchanocoDirector) {
	proxy.Directors = append(proxy.Directors, director)
}

func (proxy *OchanocoProxy) AddModifyResponse(modifyResponse OchanocoResponse) {
	proxy.ModifyResponses = append(proxy.ModifyResponses, modifyResponse)
}

func NewOchancoProxy(
	directors []OchanocoDirector,
	modifyResponses []OchanocoResponse,
	database *Database,
) OchanocoProxy {
	proxy := OchanocoProxy{}

	proxy.Directors = directors
	proxy.ModifyResponses = modifyResponses
	proxy.Database = database

	director := func(req *http.Request) {
		proxy.Director(req)
	}

	modifyResp := func(resp *http.Response) error {
		return proxy.ModifyResponse(resp)
	}

	proxy.ReverseProxy = &httputil.ReverseProxy{
		Director:       director,
		ModifyResponse: modifyResp,
	}

	return proxy
}
