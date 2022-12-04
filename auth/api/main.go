package main

import (
	"gin_line_login"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	lineLogin, err := gin_line_login.DefaultLineLogin(r)
	if err != nil {
		panic(err)
	}

	r.GET("/", lineLogin.AuthMiddleware(), func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "loggined!"})
	})

	r.GET("/login", func(c *gin.Context) {
		lineLogin.Login(c)
	})

	r.GET("/logout", func(c *gin.Context) {
		lineLogin.Logout(c)
		c.JSON(200, gin.H{"message": "logout"})
	})

	r.GET("/unauthorized", func(c *gin.Context) {
		c.JSON(401, gin.H{"message": "unauthorized"})
	})
	r.Run()
}
