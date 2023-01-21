package main

import (
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func AuthServer(secret string) *gin.Engine {
	r := gin.Default()

	store := cookie.NewStore([]byte(secret))
	r.Use(sessions.Sessions(string(secret), store))

	LineLoginFunctionalPoints(r)
	LineLoginFrontPoints(r)

	return r
}

func ProxyServer() *gin.Engine {
	r := gin.Default()

	directors := DEFAULT_DIRECTORS
	modifyResponses := DEFAULT_MODIFY_RESPONSES

	db, err := InitDB(DB_CONFIG)
	if err != nil {
		log.Fatalf("failed to init db: %v", err)
	}

	proxy := NewOchancoProxy(r, directors, modifyResponses, db)

	return proxy.Engine

}

func main() {
	secret := "secret"

	authServ := AuthServer(secret)
	proxyServ := ProxyServer()

	go authServ.Run(":8080")
	proxyServ.Run(":9000")

}
