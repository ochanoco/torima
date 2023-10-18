package core

import (
	gin_ninsho "github.com/ochanoco/ninsho/extension/gin"
)

/* configuration of DB */
var DB_TYPE = readEnv("OCHANOCO_DB_TYPE", "sqlite3")
var DB_CONFIG = readEnv("OCHANOCO_DB_CONFIG", "file:./data/db.sqlite3?_fk=1")
var SECRET = readEnv("OCHANOCO_SECRET", randomString(32))

/* other */
var DEFAULT_DIRECTORS = []OchanocoDirector{
	AuthDirector,
	DefaultRouteDirector,
	ThirdPartyDirector,
	LogDirector,
}

var DEFAULT_MODIFY_RESPONSES = []OchanocoModifyResponse{
	InjectServiceWorkerModifyResponse,
}

var DEFAULT_PROXYWEB_PAGES = []OchanocoProxyWebPage{
	ConfigWeb,
	StaticWeb,
	LoginWebs,
}

var CONFIG_FILE = "./config.yaml"
var STATIC_FOLDER = "./static"

var AUTH_PATH = gin_ninsho.NinshoGinPath{
	Unauthorized: "/auth/login",
	Callback:     "/auth/callback",
	AfterAuth:    "/_ochanoco/back",
}

var CLIENT_ID = readEnvOrPanic("OCHANOCO_CLIENT_ID")
var CLIENT_SECRET = readEnvOrPanic("OCHANOCO_CLIENT_SECRET")
