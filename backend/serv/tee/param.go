package tee

import (
	"github.com/ochanoco/proxy/core"
	"github.com/ochanoco/proxy/serv/line"

)

var DEFAULT_PROXYWEB_PAGES = []core.OchanocoProxyWebPage{
	line.ProxyDirectLineLogin,
}
