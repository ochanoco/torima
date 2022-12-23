package proxy

import (
	"fmt"
	"gin_line_login"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func InitIdPWeb(r *gin.Engine) {
	lineLogin, err := gin_line_login.NewLineLogin(r, "/login", "/auth/callback", "/redirect")
	if err != nil {
		panic(err)
	}

	r.GET("/auth/", lineLogin.AuthMiddleware(), func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "loggined!"})
	})

	r.GET("/auth/login", func(c *gin.Context) {
		lineLogin.Login(c)
	})

	r.GET("/auth/logout", func(c *gin.Context) {
		lineLogin.Logout(c)
		c.JSON(200, gin.H{"message": "logout"})
	})

	r.GET("/login", func(c *gin.Context) {
		// todo: authenticate servicer
		clientId, err := c.Get("client_id")
		if err {
			return
		}

		session := sessions.Default(c)
		session.Set("client_id", clientId)
		session.Save()

		c.HTML(http.StatusOK, "login.html", gin.H{})
	})

	r.GET("/redirect", func(c *gin.Context) {
		const redirect_uri = "http://localhost:8080/ochanoco/callback"
		c.Redirect(http.StatusTemporaryRedirect, redirect_uri)
	})

}

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