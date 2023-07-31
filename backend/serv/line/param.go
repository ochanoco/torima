package line

import "github.com/ochanoco/proxy/core"

var DEFAULT_PROXYWEB_PAGES = []core.OchanocoProxyWebPage{
	ProxyDirectLineLogin,
	core.IgnoreListWeb,
	core.StaticWeb,
}

var DEFAULT_DIRECTORS = []core.OchanocoDirector{
	core.AuthDirector,
	core.DefaultRouteDirector,
	core.ThirdPartyDirector,
	AttachUserIdDirector,
	core.LogDirector,
}
