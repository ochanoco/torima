package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func AuthServer() {
	secret := []byte("secret")

	r := gin.Default()
	r.LoadHTMLGlob(BASE_NEXT_PATH + "/*.html")
	r.Static("/_next/", BASE_NEXT_PATH+"_next/")

	store := cookie.NewStore(secret)
	r.Use(sessions.Sessions("mysession", store))

	InitProxyWeb(r)
	InitIdPWeb(r)

	r.Run()
}

func ProxyServer() {
	directors := []OchanocoDirector{}
	modifyResponses := []OchanocoModifyResponse{}

	db, err := InitDB(DB_CONFIG)
	if err != nil {
		log.Fatalf("failed to init db: %v", err)
	}

	proxy := NewOchancoProxy(directors, modifyResponses, db)
	server := http.Server{
		Addr:    ":9000",
		Handler: proxy.ReverseProxy,
	}

	server.ListenAndServe()
}

func main() {
	go AuthServer()
	ProxyServer()
}
