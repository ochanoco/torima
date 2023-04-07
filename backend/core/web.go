package core

import (
	"fmt"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func SimpleReverseProxy(name string, baseUrl *url.URL, r *gin.RouterGroup) {
	SetupParsingUrl()

	fmt.Printf("test2: %v\n", baseUrl)

	r.GET(name, func(c *gin.Context) {
		fmt.Printf("test1: %v\n", baseUrl)
		proxy := httputil.NewSingleHostReverseProxy(baseUrl)
		proxy.ServeHTTP(c.Writer, c.Request)
	})
}

func SimpleReverseProxies(pageList []string, baseUrl *url.URL, r *gin.RouterGroup) {
	fmt.Printf("test3: %v\n", baseUrl)

	for _, value := range pageList {
		SimpleReverseProxy(value, baseUrl, r)
	}
}

func NextJSProxyPage(baseUrl *url.URL, r *gin.RouterGroup) {
	fmt.Printf("test4: %v\n", baseUrl)

	nextJsProxyPaths := []string{
		"/_next/webpack-hmr",
		"/_next/static/chunks/:file",
		"/_next/static/chunks/pages/:file",
		"/_next/static/development/:file",
	}

	SimpleReverseProxies(nextJsProxyPaths, baseUrl, r)
}

// func DeriveOchanocoNextJSProxyPage(baseUrl *url.URL) func(proxy *OchanocoProxy, r *gin.RouterGroup) {
// 	return func(proxy *OchanocoProxy, r *gin.RouterGroup) {
// 		NextJSProxyPage(baseUrl, r)
// 	}
// }

// func DeriveOchanocoProxyPage(name string, baseUrl *url.URL) func(proxy *OchanocoProxy, r *gin.RouterGroup) {
// 	return func(proxy *OchanocoProxy, r *gin.RouterGroup) {
// 		SimpleReverseProxy(name, baseUrl, r)
// 	}
// }

func AuthSiteLoginPage(proxy *OchanocoProxy, r *gin.RouterGroup) {
	SimpleReverseProxy("/login", AuthWebBaseUrl, r)
}

func AuthSiteNextJS(proxy *OchanocoProxy, r *gin.RouterGroup) {
	NextJSProxyPage(AuthWebBaseUrl, r)
}
