package main

import (
	"context"
	"errors"
	"log"

	"github.com/tracer-silver-bullet/tracer-silver-bullet/proxy/ent"
)

type Database struct {
	ctx    context.Context
	client *ent.Client
}

var db *Database

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

func savePageOnProj(db *Database, projc *ent.Project, wlc *ent.WhiteListCreate) (*ent.Project, error) {
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
