package database

import (
	"context"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/ochanoco/database/ent"
)

var DB_TYPE = "sqlite3"
var DB_CONFIG = "file:./db.sqlite3?_fk=1"
var WHITELIST_FILE = "./whitelist.json"

var LOGIN_REDIRECT_PAGE_URL = "http://localhost:3000/redirect"
var ERROR_PAGE_URL = "http://localhost:3000/error"

var AUTH_PAGE_DOMAIN = "localhost:9000"
var AUTH_PAGE_DESTINATION = "localhost:3000"

type Database struct {
	ctx    context.Context
	client *ent.Client
}

func InitDB(path string) (*Database, error) {
	client, err := ent.Open(DB_TYPE, DB_CONFIG)

	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	ctx := context.Background()

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	dbl := new(Database)
	dbl.ctx = ctx
	dbl.client = client

	return dbl, err
}
