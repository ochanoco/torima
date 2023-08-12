package core

import (
	"github.com/gin-gonic/gin"
)

func StaticWeb(proxy *OchanocoProxy, r *gin.RouterGroup) {
	r.Use(func() gin.HandlerFunc {
		return func(c *gin.Context) {
			c.Writer.Header().Set("Service-Worker-Allowed", "/")
		}
	}())

	r.Static("/static", STATIC_FOLDER)
}

func ConfigWeb(proxy *OchanocoProxy, r *gin.RouterGroup) {
	r.GET("/config.json", func(c *gin.Context) {
		c.JSON(200, proxy.Config.ProtectionScope)
	})
}
