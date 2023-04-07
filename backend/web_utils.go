package main

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func deriveSimpelProxyFunc(baseUrl *url.URL) func(c *gin.Context) {
	return func(c *gin.Context) {
		proxy := httputil.NewSingleHostReverseProxy(baseUrl)
		proxy.ServeHTTP(c.Writer, c.Request)

	}
}

func SimpleProxyPage(pageList []string, baseUrl *url.URL, r *gin.RouterGroup) {
	for _, value := range pageList {
		r.GET(value, deriveSimpelProxyFunc(baseUrl))
	}
}

func NextJSProxyPage[T gin.RouterGroup](baseUrl *url.URL, r *gin.RouterGroup) {
	nextJsProxyPaths := []string{
		"/_next/webpack-hmr",
		"/_next/static/chunks/:file",
		"/_next/static/chunks/pages/:file",
		"/_next/static/development/:file",
	}

	SimpleProxyPage(nextJsProxyPaths, baseUrl, r)
}
