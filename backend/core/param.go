package core

import "os"

/* configuration of URLs */
var PROTOCOL = os.Getenv("OCHANOCO_PROTOCOL")

// http://localhost:8080/
var PROXY_PORT = os.Getenv("OCHANOCO_PROXY_PORT")
var PROXY_HOST = os.Getenv("OCHANOCO_PROXY_HOST")
var PROXY_BASE = PROTOCOL + PROXY_HOST
var PROXY_CALLBACK_URL = PROXY_BASE + "/ochanoco/callback"
var PROXY_LOGIN_URL = PROXY_BASE + "/ochanoco/login"
var PROXY_REDIRECT_URL = PROXY_BASE + "/ochanoco/redirect"

// https://localhost:8081
var AUTH_PORT = os.Getenv("OCHANOCO_AUTH_PORT")
var AUTH_HOST = os.Getenv("OCHANOCO_AUTH_HOST")
var AUTH_BASE = PROTOCOL + AUTH_HOST

// https://localhost:3000
var PROXYWEB_PORT = os.Getenv("OCHANOCO_PROXYWEB_PORT")
var PROXYWEB_HOST = os.Getenv("OCHANOCO_PROXYWEB_HOST")
var PROXYWEB_BASE = PROTOCOL + PROXYWEB_HOST
var ERROR_URL = PROXYWEB_BASE + "/error"

// https://localhost:3000
var AUTHWEB_PORT = os.Getenv("OCHANOCO_AUTHWEB_PORT")
var AUTHWEB_HOST = os.Getenv("OCHANOCO_AUTHWEB_HOST")
var AUTHWEB_BASE = PROTOCOL + AUTHWEB_HOST

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

var ADD_USER_ID = false
var CONFIG_FILE = "./config.yaml"
