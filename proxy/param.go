package proxy

import (
	"github.com/ochanoco/torima/core"
	"github.com/ochanoco/torima/extension"
	"github.com/ochanoco/torima/extension/directors"
)

/* other */
var DEFAULT_DIRECTORS = core.TorimaDirectors{
	directors.BeforeLogDirector,
	directors.SanitizeHeaderDirector,
	directors.SkipAuthDirector,
	directors.ForceAuthDirector,
	directors.AuthDirector,
	directors.DefaultRouteDirector,
	directors.ThirdPartyDirector,
	directors.AfterLogDirector,
}

var DEFAULT_MODIFY_RESPONSES = core.TorimaModifyResponses{
	extension.InjectServiceWorkerModifyResponse,
}

var DEFAULT_PROXYWEB_PAGES = []core.TorimaProxyWebPage{
	extension.ConfigWeb,
	extension.StaticWeb,
	extension.LoginWebs,
}
