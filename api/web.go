package main

import (
	"fmt"
	"gin_line_login"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LineLoginFunctionalPoints(r *gin.Engine) {
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
		const redirect_uri = "http://localhost:8080/ochanoco/callback"
		c.Redirect(http.StatusTemporaryRedirect, redirect_uri)
	})
}

func LineLoginFrontPoints(r *gin.Engine) {
	proxyFunc := func(c *gin.Context) {
		url, err := url.Parse(OCHANOCO_FRONT_LOGIN_DOMAIN)
		if err != nil {
			panic(err)
		}
		proxy := httputil.NewSingleHostReverseProxy(url)

		proxy.Director = func(req *http.Request) {
			req.Header = c.Request.Header
			req.Host = url.Host
			req.URL.Scheme = url.Scheme
			req.URL.Host = url.Host

			fmt.Printf("host: %v\n", url.Host)
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	}

	proxyToPageFunc := func(c *gin.Context) {
		// todo: authenticate servicer
		clientId, isExists := c.Get("client_id")
		if !isExists {
			panic("client_id is not found")
		}

		session := sessions.Default(c)
		session.Set("client_id", clientId)
		session.Save()

		proxyFunc(c)
	}

	paths := []string{
		"/_next/webpack-hmr",
		"/_next/static/chunks/:file",
		"/_next/static/chunks/pages/:file",
		"/_next/static/development/:file",
	}

	for _, value := range paths {
		r.GET(value, proxyFunc)
	}

	r.GET("/login", proxyToPageFunc)
}
