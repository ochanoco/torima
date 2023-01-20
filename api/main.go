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

	store := cookie.NewStore(secret)
	r.Use(sessions.Sessions("mysession", store))

	InitIdPWeb(r)

	r.Run()
}

func ProxyServer() {
	directors := []OchanocoDirector{
		MainDirector,
	}

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
