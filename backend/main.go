package main

import (
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func AuthServer(secret string, proxy *OchanocoProxy) *gin.Engine {
	r := gin.Default()

	store := cookie.NewStore([]byte(secret))
	r.Use(sessions.Sessions("session", store))

	LineLoginFunctionalPoints(r, proxy)
	LineLoginFrontPoints(r, proxy)

	return r
}

func ProxyServer(secret string) *OchanocoProxy {
	r := gin.Default()

	store := cookie.NewStore([]byte(secret))
	r.Use(sessions.Sessions("ochanoco-session", store))

	directors := DEFAULT_DIRECTORS
	modifyResponses := DEFAULT_MODIFY_RESPONSES

	db, err := InitDB(DB_CONFIG)
	if err != nil {
		log.Fatalf("failed to init db: %v", err)
	}

	proxy := NewOchancoProxy(r, directors, modifyResponses, db)

	return &proxy
}

func setup() {
	setupParsingUrl()
}

func main() {
	secret := "testest"

	setup()

	proxyServ := ProxyServer(secret)
	authServ := AuthServer(secret, proxyServ)

	go authServ.Run(AUTH_PORT)
	proxyServ.Engine.Run(PROXY_PORT)

}
