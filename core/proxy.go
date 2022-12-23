package main

import (
	"net/http"
	"net/http/httputil"
)

type OchanocoProxy struct {
	Directors       []func(req *http.Request)
	ModifyResponses []func(req *http.Response)
	ReverseProxy    *httputil.ReverseProxy
	Server          *http.Server
}

/**
 * Directors is a list of functions that modify the
 * request before it is sent to the target server.
 **/
func (proxy *OchanocoProxy) Director(req *http.Request) {
	for _, d := range proxy.Directors {
		d(req)
	}
}

/**
  * ModifyResponses is a list of functions that modify the
  * response before it is sent to the client.
**/
func (proxy *OchanocoProxy) ModifyResponse(res *http.Response) error {
	for _, mR := range proxy.ModifyResponses {
		mR(res)
	}

	return nil
}

func NewOchancoProxy(
	Directors []func(req *http.Request),
	ModifyResponses []func(req *http.Response),
) OchanocoProxy {
	proxy := OchanocoProxy{}

	proxy.Directors = Directors
	proxy.ModifyResponses = ModifyResponses

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

	proxy.Server = &http.Server{
		Addr:    ":9000",
		Handler: proxy.ReverseProxy,
	}

	return proxy
}
