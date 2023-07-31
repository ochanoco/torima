package core

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func SimpleReverseProxy(name string, baseUrl *url.URL, r *gin.RouterGroup) {

	r.GET(name, func(c *gin.Context) {
		proxy := httputil.NewSingleHostReverseProxy(baseUrl)
		proxy.ServeHTTP(c.Writer, c.Request)
	})
}

func SimpleReverseProxies(pageList []string, baseUrl *url.URL, r *gin.RouterGroup) {
	for _, value := range pageList {
		SimpleReverseProxy(value, baseUrl, r)
	}
}

func NextJSProxyPage(baseUrl *url.URL, r *gin.RouterGroup) {
	nextJsProxyPaths := []string{
		"/_next/webpack-hmr",
		"/_next/static/chunks/:file",
		"/_next/static/chunks/pages/:file",
		"/_next/static/development/:file",
	}

	SimpleReverseProxies(nextJsProxyPaths, baseUrl, r)
}

func AuthSiteLoginPage(proxy *OchanocoProxy, r *gin.RouterGroup) {
	SimpleReverseProxy("/login", AuthWebBaseUrl, r)
}

func AuthSiteNextJS(proxy *OchanocoProxy, r *gin.RouterGroup) {
	NextJSProxyPage(AuthWebBaseUrl, r)
}
