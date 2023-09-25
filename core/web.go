package core

import (
	"github.com/gin-contrib/sessions"
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
	r.GET("/status", func(c *gin.Context) {
		session := sessions.Default(c)
		userId := session.Get("userId")

		c.JSON(200, gin.H{
			"protection_scope": proxy.Config.ProtectionScope,
			"white_list_path":  proxy.Config.WhiteListPath,
			"is_authenticated": userId != nil, // is it needed?.
		})
	})
}
