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

	LineLoginFunctionalPoints(proxy, &r.RouterGroup)
	LineLoginFrontPoints(r, proxy)

	return r
}

func ProxyServer(secret string) *OchanocoProxy {
	r := gin.Default()

	store := cookie.NewStore([]byte(secret))
	r.Use(sessions.Sessions("ochanoco-session", store))

	db, err := InitDB(DB_CONFIG)
	if err != nil {
		log.Fatalf("failed to init db: %v", err)
	}

	proxy := NewOchancoProxy(r, DEFAULT_DIRECTORS, DEFAULT_MODIFY_RESPONSES, DEFAULT_PROXYWEB_PAGES, db)

	return &proxy
}

func main() {
	secret := "testest"

	setupParsingUrl()

	proxyServ := ProxyServer(secret)
	authServ := AuthServer(secret, proxyServ)

	go authServ.Run(AUTH_PORT)
	proxyServ.Engine.Run(PROXY_PORT)

}
