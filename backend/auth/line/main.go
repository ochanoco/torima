package line

import (
	"github.com/ochanoco/proxy/core"
)

const NAME = "line"

func Run() *core.OchanocoProxy {
	secret := "testest"
	core.DEFAULT_PROXYWEB_PAGES = DEFAULT_PROXYWEB_PAGES

	core.SetupParsingUrl()

	proxyServ := core.ProxyServer(secret)
	return proxyServ
}

func Main() {
	proxyServ := Run()
	proxyServ.Engine.Run(core.PROXY_PORT)
}
