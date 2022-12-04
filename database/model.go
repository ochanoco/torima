package database

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

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

var DB *Database

func InitDB() (*Database, error) {
	err := errors.New("error")
	dbl := new(Database)

	client, err := ent.Open(DB_TYPE, DB_CONFIG)

	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	ctx := context.Background()

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	dbl = new(Database)
	dbl.ctx = ctx
	dbl.client = client

	DB = dbl

	return dbl, err
}

func MigrateWhiteList() error {
	var urls []string

	b, _ := os.ReadFile(WHITELIST_FILE)
	err := json.Unmarshal(b, &urls)

	if err != nil {
		log.Fatalf("failed to load migrate.json: %v", err)
		return err
	}

	projc := CreateServiceProvider(DB, AUTH_PAGE_DOMAIN, AUTH_PAGE_DESTINATION)
	proj, nil := projc.Save(DB.ctx)

	if err != nil {
		return fmt.Errorf("failed creating project: %v", err)
	}

	for _, url := range urls {
		wl := CreateWhiteList(DB, url)
		proj, err = SaveWhiteListOnProj(DB, proj, wl)

		if err != nil {
			return fmt.Errorf("failed add white list to project: %v", err)
		}
	}

	return nil
}

func CreateWhiteList(db *Database, path string) *ent.WhiteListCreate {
	wl := db.client.WhiteList.
		Create().
		SetPath(path)

	return wl
}

func CreateServiceProvider(db *Database, host string, destinationIP string) *ent.ServiceProviderCreate {
	proj := db.client.ServiceProvider.
		Create().
		SetHost(host).
		SetDestinationIP(destinationIP)

	return proj
}

func SaveWhiteListOnProj(db *Database, projc *ent.ServiceProvider, wlc *ent.WhiteListCreate) (*ent.ServiceProvider, error) {
	wl, err := wlc.Save(db.ctx)

	if err != nil {
		return projc, err
	}

	proj, err := projc.
		Update().
		AddWhitelists(wl).
		Save(db.ctx)

	return proj, err
}
