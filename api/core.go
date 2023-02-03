package main

import (
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

type OchanocoDirector = func(proxy *OchanocoProxy, req *http.Request, c *gin.Context) bool
type OchanocoModifyResponse = func(proxy *OchanocoProxy, req *http.Response, c *gin.Context) bool

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

	proxy.Engine = r

	r.NoRoute(func(c *gin.Context) {
		director := func(req *http.Request) {
			proxy.Director(req, c)
		}

		modifyResp := func(resp *http.Response) error {
			return proxy.ModifyResponse(resp, c)
		}

		proxy := httputil.ReverseProxy{
			Director:       director,
			ModifyResponse: modifyResp,
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	})

	return proxy
}
