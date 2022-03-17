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

func init_db() (*Database, error) {
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
