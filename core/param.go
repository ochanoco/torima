package core

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
