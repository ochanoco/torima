package core

import (
	"context"
	"fmt"
	"log"
	"time"

	_ "sqlite3"

	"github.com/ochanoco/proxy/ent"
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

func (db *Database) CreateServiceLog(t time.Time, header string, body []byte) *ent.ServiceLogCreate {
	sl := db.Client.ServiceLog.
		Create().
		SetTime(t).
		SetHeaders(header).
		SetBody(body)

	return sl
}

func (db *Database) SaveServiceLog(l *ent.ServiceLogCreate) (*ent.ServiceLog, error) {
	code, err := l.Save(db.Ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to save save log: %v", err)
	}

	return code, err
}