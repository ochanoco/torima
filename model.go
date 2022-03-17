package main

import (
	"context"
	"errors"
	"log"

	"github.com/tracer-silver-bullet/tracer-silver-bullet/proxy/ent"
)

var DB_TYPE = "sqlite3"

var DB_CONFIG = "file:./db.sqlite3?_fk=1"

type Database struct {
	ctx    context.Context
	client *ent.Client
}

func initDB() (*Database, error) {
	err := errors.New("error")
	db := new(Database)

	client, err := ent.Open(DB_TYPE, DB_CONFIG)

	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
		return db, err
	}

	ctx := context.Background()

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
		return db, err
	}

	db = new(Database)
	db.ctx = ctx
	db.client = client

	return db, err
}

func createAndSavePage(db *Database, url string, skip bool) (*ent.Page, error) {
	page, err := db.client.Page.
		Create().
		SetURL(url).
		SetSkip(skip).
		Save(db.ctx)

	return page, err
}

func createAndSaveProj(db *Database, lineId string, name string) (*ent.Project, error) {
	page, err := db.client.Project.
		Create().
		SetLineID(lineId).
		SetName(name).
		Save(db.ctx)

	return page, err
}

func addPageToProj(db *Database, proj *ent.Project, page *ent.Page) (*ent.Project, error) {
	proj, err := proj.Update().
		AddPages(page).
		Save(db.ctx)

	return proj, err
}
