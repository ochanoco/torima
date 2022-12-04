package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitProxyWeb(r *gin.Engine) {
	r.GET("/ochanoco/redirect", func(c *gin.Context) {
		c.HTML(http.StatusOK, "redirect.html", gin.H{})
	})

	r.GET("/ochanoco/_redirect", func(c *gin.Context) {
		// todo: analyis risk of forge host
		clientId := c.Request.Host

		url := fmt.Sprintf("%v?clientId=%v", OCHANOCO_LOGIN_URL, clientId)
		c.Redirect(http.StatusTemporaryRedirect, url)
	})

	r.GET("/ochanoco/callback", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "redirected from ochanoco"})
	})
}
