package core

import "os"

/* configuration of DB */
var DB_TYPE = os.Getenv("OCHANOCO_DB_TYPE")
var DB_CONFIG = os.Getenv("OCHANOCO_DB_CONFIG")

/* other */
var DEFAULT_DIRECTORS = []OchanocoDirector{
	AuthDirector,
	DefaultRouteDirector,
	ThirdPartyDirector,
	LogDirector,
}

var DEFAULT_MODIFY_RESPONSES = []OchanocoModifyResponse{
	LogModifyResponse,
}

var DEFAULT_PROXYWEB_PAGES = []OchanocoProxyWebPage{
	IgnoreListWeb,
	StaticWeb,
}

var DEFAULT_ERROR_HANDLER = errorMiddleware()

var CONFIG_FILE = "./config.yaml"
var STATIC_FOLDER = "../static"
