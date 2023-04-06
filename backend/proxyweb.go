package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ProxyWebAuthPages(proxy *OchanocoProxy, r *gin.RouterGroup) {
	r.GET("/redirect", func(c *gin.Context) {
		callback_path := c.Query("callback_path")

		if callback_path == "" {
			callback_path = "/"
		}

		session := sessions.Default(c)
		session.Set("callback_path", callback_path)

		c.JSON(200, gin.H{"message": "hi"})

	})

	r.GET("/callback", func(c *gin.Context) {
		token := c.Query("authorization_code")
		if token == "" {
			panic("failed to get authorization_code")
		}

		session := sessions.Default(c)
		// session.Set("authorization_code", token)

		// if session.Save() != nil {
		// 	panic("failed to save authorization_code")
		// }

		callbackPath := session.Get("callback_path")
		if callbackPath == nil {
			callbackPath = "/"
		}

		// req.URL.Path = callbackPath.(string)
		// req.URL.Host = req.Host
		// req.URL.RawQuery = ""

		// return FINISHED

		c.JSON(200, gin.H{"message": "logout"})
	})

}

func ProxyLoginRedirectPage(proxy *OchanocoProxy, r *gin.RouterGroup) {
	r.GET("/:name",
		deriveSimpelProxyFunc(ProxyWebBaseUrl))
	NextJSProxyPage(ProxyWebBaseUrl, r)
}
