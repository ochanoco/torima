package core

import (
	gin_ninsho "github.com/ochanoco/ninsho/extension/gin"
)

/* configuration of DB */
var DB_TYPE = readEnv("TORIMA_DB_TYPE", "sqlite3")
var DB_CONFIG = readEnv("TORIMA_DB_CONFIG", "file:./data/db.sqlite3?_fk=1")
var SECRET = readEnv("TORIMA_SECRET", randomString(32))

/* other */
var DEFAULT_DIRECTORS = TorimaDirectors{
	BeforeLogDirector,
	SanitizeHeaderDirector,
	SkipAuthDirector,
	AuthDirector,
	DefaultRouteDirector,
	ThirdPartyDirector,
	AfterLogDirector,
}

var DEFAULT_MODIFY_RESPONSES = TorimaModifyResponses{
	InjectServiceWorkerModifyResponse,
}

var DEFAULT_PROXYWEB_PAGES = []TorimaProxyWebPage{
	ConfigWeb,
	StaticWeb,
	LoginWebs,
}

var CONFIG_FILE = "./config.yaml"
var STATIC_FOLDER = "./static"

var AUTH_PATH = gin_ninsho.NinshoGinPath{
	Unauthorized: "/auth/login",
	Callback:     "/auth/callback",
	AfterAuth:    "/_torima/back",
}

var CLIENT_ID = readEnvOrPanic("TORIMA_CLIENT_ID")
var CLIENT_SECRET = readEnvOrPanic("TORIMA_CLIENT_SECRET")
