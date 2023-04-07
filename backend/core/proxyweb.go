package core

import "github.com/gin-gonic/gin"

func ProxyLoginRedirectPage(proxy *OchanocoProxy, r *gin.RouterGroup) {
	r.GET("/:name",
		DeriveSimpelProxyFunc(ProxyWebBaseUrl))
	NextJSProxyPage(ProxyWebBaseUrl, r)
}
