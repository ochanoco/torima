package cloud

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/ochanoco/proxy/core"
)

const NAME = "cloud"

func AuthServer(secret string, proxy *core.OchanocoProxy) *gin.Engine {
	r := gin.Default()

	store := cookie.NewStore([]byte(secret))
	r.Use(sessions.Sessions("session", store))

	CloudLoginFunctionalPoints(proxy, &r.RouterGroup)
	CloudLoginFrontPoints(r, proxy)

	return r
}

func Main() {
	secret := "testest"

	core.SetupParsingUrl()

	proxyServ := core.ProxyServer(secret)
	authServ := AuthServer(secret, proxyServ)

	go authServ.Run(core.AUTH_PORT)
	proxyServ.Engine.Run(core.PROXY_PORT)

}
