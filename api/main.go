package main

import (
	"gin_line_login"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const BASE_NEXT_PATH = "../app/out/"

func main() {
	r := gin.Default()
	r.LoadHTMLGlob(BASE_NEXT_PATH + "/*.html")
	r.Static("/_next/", BASE_NEXT_PATH+"_next/")

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

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

	r.GET("/ochanoco/redirect", func(c *gin.Context) {
		c.HTML(http.StatusOK, "redirect.html", gin.H{})
	})

	r.GET("/ochanoco/callback", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "redirect"})
	})

	r.GET("/login", func(c *gin.Context) {
		// todo: authenticate servicer
		_, err := c.Get("client_id")
		if err {
		}

		_, err = c.Get("client_secret")
		if err {

		}

		c.HTML(http.StatusOK, "login.html", gin.H{})
	})

	r.Run()
}
