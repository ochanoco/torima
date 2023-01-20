package main

var LOGIN_REDIRECT_PAGE_URL = "http://localhost:8080/redirect"
var ERROR_PAGE_URL = "http://localhost:8080/error"
var PROXY_REDIRECT_URL = "http://localhost:8080/"

var AUTH_PAGE_DOMAIN = "localhost:8080"
var AUTH_PAGE_DESTINATION = "localhost:8080/ochanoco/callback"

var PROXYWEB_DOMAIN = "localhost:3000"

var DB_TYPE = "sqlite3"
var DB_CONFIG = "file:./db.sqlite3?_fk=1"
var WHITELIST_FILE = "./whitelist.json"

const BASE_NEXT_PATH = "../app/out/"
const OCHANOCO_LOGIN_URL = "http://localhost:8080/login"

const TEST_DB_PATH = "file:ent?mode=memory&cache=shared&_fk=1"
