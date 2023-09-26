package core

import (
	"gin_line_login"

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

func LoginWebs(proxy *OchanocoProxy, r *gin.RouterGroup) {
	lineLogin, err := gin_line_login.NewLineLoginWithEnvironment(r, "/ochanoco/auth/login", "/auth/callback", "/_ochanoco/back")
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
