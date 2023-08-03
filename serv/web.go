package serv

import (
	"gin_line_login"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/ochanoco/proxy/core"
)

func ProxyDirectLineLogin(proxy *core.OchanocoProxy, r *gin.RouterGroup) {
	lineLogin, err := gin_line_login.NewLineLoginWithEnvironment(r, "/ochanoco/auth/login", "/auth/callback", "/ochanoco/auth/redirect")
	if err != nil {
		panic(err)
	}

	r.GET("/login", func(c *gin.Context) {
		lineLogin.Login(c)
	})

	r.GET("/auth/logout", func(c *gin.Context) {
		lineLogin.Logout(c)
		c.JSON(200, gin.H{"message": "logout"})
	})

	r.GET("/auth/status", lineLogin.AuthMiddleware(), func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "loggined!"})
	})

	r.GET("/auth/redirect", lineLogin.AuthMiddleware(), func(c *gin.Context) {
		session := sessions.Default(c)
		userId := session.Get("userId")

		c.JSON(200, gin.H{"user_id": userId})
	})
}
