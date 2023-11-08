package core

import (
	"context"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/ochanoco/torima/ent"
)

type Database struct {
	Ctx    context.Context
	Client *ent.Client
}

func InitDB(path string) (*Database, error) {
	client, err := ent.Open(DB_TYPE, path)

	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	ctx := context.Background()

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	dbl := new(Database)
	dbl.Ctx = ctx
	dbl.Client = client

	return dbl, err
}

func (db *Database) CreateRequestLog(header string, body []byte, flag string) *ent.RequestLogCreate {
	t := time.Now()

	sl := db.Client.RequestLog.
		Create().
		SetTime(t).
		SetHeaders(header).
		SetBody(body).
		SetFlag(flag)

	return sl
}
