package line

import "github.com/ochanoco/proxy/core"

var DEFAULT_PROXYWEB_PAGES = []core.OchanocoProxyWebPage{
	ProxyDirectLineLogin,
	core.ProxyLoginRedirectPage,
}
