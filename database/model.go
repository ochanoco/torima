package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/ochanoco/database/proxy/ent"
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

var db *Database

func initDB() (*Database, error) {
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

	db = dbl

	return dbl, err
}

func migrateWhiteList() error {
	var urls []string

	b, _ := os.ReadFile(WHITELIST_FILE)
	err := json.Unmarshal(b, &urls)

	if err != nil {
		log.Fatalf("failed to load migrate.json: %v", err)
		return err
	}

	projc := createProject(db, AUTH_PAGE_DOMAIN, AUTH_PAGE_DESTINATION, "root", "root")
	proj, nil := projc.Save(db.ctx)

	if err != nil {
		fmt.Errorf("failed creating project: %v", err)
		return err
	}

	for _, url := range urls {
		wl := createWhiteList(db, url)
		proj, err = saveWhiteListOnProj(db, proj, wl)

		if err != nil {
			fmt.Errorf("failed add white list to project: %v", err)
			return err
		}
	}

	return nil
}

func createWhiteList(db *Database, url string) *ent.WhiteListCreate {
	wl := db.client.WhiteList.
		Create().
		SetURL(url)

	return wl
}

func createProject(db *Database, domain string, destination string, lineId string, name string) *ent.ProjectCreate {
	proj := db.client.Project.
		Create().
		SetDomain(domain).
		SetDestination(destination).
		SetLineID(lineId).
		SetName(name)

	return proj
}

func saveWhiteListOnProj(db *Database, projc *ent.Project, wlc *ent.WhiteListCreate) (*ent.Project, error) {
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
