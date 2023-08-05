package core

import (
	"github.com/gin-gonic/gin"
)

func StaticWeb(proxy *OchanocoProxy, r *gin.RouterGroup) {
	r.Static("/static", STATIC_FOLDER)
}

func IgnoreListWeb(proxy *OchanocoProxy, r *gin.RouterGroup) {
	r.GET("/ignores.json", func(c *gin.Context) {
		c.JSON(200, proxy.Config.IgnoredOrigins)
	})
}
