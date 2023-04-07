package main

import (
	"fmt"
	"gin_line_login"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LineLoginFunctionalPoints(proxy *OchanocoProxy, r *gin.RouterGroup) {
	lineLogin, err := gin_line_login.NewLineLoginWithEnvironment(r, "/login", "/auth/callback", "/redirect")
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

	r.GET("/redirect", func(c *gin.Context) {
		session := sessions.Default(c)
		host := session.Get("host")

		codeCreate, err := proxy.Database.CreateRandomAuthorizationCode()

		if err != nil {
			panic("failed to generate authorization code")
		}

		code, err := proxy.Database.SaveAuthorizationCode(codeCreate)

		if err != nil {
			panic("failed to save authorization code")
		}

		redirect_uri := fmt.Sprintf("http://%v/ochanoco/callback?authorization_code=%v", host, code.Token)
		c.Redirect(http.StatusTemporaryRedirect, redirect_uri)
	})
}

func LineLoginFrontPoints(r *gin.Engine, proxy *OchanocoProxy) {
	proxyToPageFunc := func(c *gin.Context) {
		// todo: authenticate servicer
		clientId, isExists := c.GetQuery("client_id")
		if !isExists {
			panic("client_id is not found on query params")
		}

		project, err := proxy.Database.FindServiceProviderByHost(clientId)
		if err != nil || project == nil {
			log.Fatalf("client_id is not found on DB(%v)", clientId)
		}

		session := sessions.Default(c)
		session.Set("host", project.Host)
		session.Save()

		deriveSimpelProxyFunc(AuthWebBaseUrl)(c)
	}

	r.GET("/login", proxyToPageFunc)
	NextJSProxyPage(AuthWebBaseUrl, &r.RouterGroup)
}
