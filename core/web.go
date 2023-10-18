package core

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/ochanoco/ninsho"
	gin_ninsho "github.com/ochanoco/ninsho/extension/gin"
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
	login, err := gin_ninsho.NewNinshoGin(r, &provider, &ninsho.LINE_LOGIN, OCHANOCO_BASE, &AUTH_PATH)
	if err != nil {
		panic(err)
	}

	r.GET("/login", func(c *gin.Context) {
		login.Login(c)
	})

	r.GET("/auth/logout", func(c *gin.Context) {
		login.Logout(c)
		c.JSON(200, gin.H{"message": "logout"})
	})

	r.GET("/auth/status", login.AuthMiddleware(), func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "loggined!"})
	})

	r.GET("/auth/redirect", login.AuthMiddleware(), func(c *gin.Context) {
		user, err := gin_ninsho.GetUser[ninsho.LINE_USER](c)

		if err != nil {
			panic(err)
		}

		c.JSON(200, gin.H{"user_id": user.Sub})
	})
}
