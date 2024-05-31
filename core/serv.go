package core

import (
	"fmt"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func ProxyServer() (*TorimaProxy, error) {
	secret := randomString(64)
	r := gin.Default()

	store := cookie.NewStore([]byte(secret))
	r.Use(sessions.Sessions("torima-session", store))

	db, err := InitDB(DB_CONFIG)
	if err != nil {
		log.Fatalf("failed to init db: %v", err)
	}

	config, err := readConfig()
	if config == nil {
		panic("failed to read config: " + err.Error())
	}

	if err != nil {
		fmt.Printf("failed to read config file, so set default parameters: %v", err)
	}

	printConfig(config)

	proxy := NewOchancoProxy(r, DEFAULT_DIRECTORS, DEFAULT_MODIFY_RESPONSES, DEFAULT_PROXYWEB_PAGES, config, db)

	return &proxy, nil
}
