package core

import (
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func ProxyServer(secret string) (*OchanocoProxy, error) {
	r := gin.Default()

	store := cookie.NewStore([]byte(secret))
	r.Use(sessions.Sessions("ochanoco-session", store))

	db, err := InitDB(DB_CONFIG)
	if err != nil {
		log.Fatalf("failed to init db: %v", err)
	}

	config, err := readConfig()
	err = makeError(err, "failed to read whitelist file")

	proxy := NewOchancoProxy(r, DEFAULT_DIRECTORS, DEFAULT_MODIFY_RESPONSES, DEFAULT_PROXYWEB_PAGES, &config, db)

	return &proxy, err
}
