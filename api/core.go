package main

import (
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

type OchanocoDirector = func(proxy *OchanocoProxy, req *http.Request) bool
type OchanocoModifyResponse = func(proxy *OchanocoProxy, req *http.Response) bool

type OchanocoProxy struct {
	Directors       []OchanocoDirector
	ModifyResponses []OchanocoModifyResponse
	Engine          *gin.Engine
	Database        *Database
}

func NewOchancoProxy(
	r *gin.Engine,
	directors []OchanocoDirector,
	modifyResponses []OchanocoModifyResponse,
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

	proxy.Engine = r

	r.NoRoute(func(c *gin.Context) {
		proxy := httputil.ReverseProxy{
			Director:       director,
			ModifyResponse: modifyResp,
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	})

	return proxy
}
