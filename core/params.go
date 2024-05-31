package core

import "github.com/ochanoco/torima/utils"

/* configuration of DB */
var DB_TYPE = utils.ReadEnv("TORIMA_DB_TYPE", "sqlite3")
var DB_CONFIG = utils.ReadEnv("TORIMA_DB_CONFIG", "file:./data/db.sqlite3?_fk=1")
var SECRET = utils.ReadEnv("TORIMA_SECRET", utils.RandomString(32))

var CONFIG_FILE = "./config.yaml"
var STATIC_FOLDER = "./static"

var CLIENT_ID = utils.ReadEnvOrPanic("TORIMA_CLIENT_ID")
var CLIENT_SECRET = utils.ReadEnvOrPanic("TORIMA_CLIENT_SECRET")
