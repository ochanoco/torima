package core

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "sqlite3"

	"github.com/ochanoco/proxy/ent"
	"github.com/ochanoco/proxy/ent/serviceprovider"
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

func (db *Database) CreateWhiteList(path string) *ent.WhiteListCreate {
	wl := db.Client.WhiteList.
		Create().
		SetPath(path)

	return wl
}

func (db *Database) CreateServiceProvider(host string, destinationIP string) *ent.ServiceProviderCreate {
	proj := db.Client.ServiceProvider.
		Create().
		SetHost(host).
		SetDestinationIP(destinationIP)

	return proj
}

func (db *Database) MigrateWhiteList() error {
	var urls []string

	b, _ := os.ReadFile(WHITELIST_PATH)
	err := json.Unmarshal(b, &urls)

	if err != nil {
		log.Fatalf("failed to load migrate.json: %v", err)
		return err
	}

	projc := db.CreateServiceProvider(AUTH_HOST, PROXY_CALLBACK_URL)
	_, err = projc.Save(db.Ctx)

	if err != nil {
		return fmt.Errorf("failed creating project: %v", err)
	}

	for _, url := range urls {
		wlc := db.CreateWhiteList(url)

		_, err := wlc.Save(db.Ctx)

		if err != nil {
			return fmt.Errorf("failed add white list to project: %v", err)
		}
	}

	return nil
}

func (db *Database) SaveServiceProvider(spc *ent.ServiceProviderCreate) (*ent.ServiceProvider, error) {
	sp, err := spc.Save(db.Ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to save service provider: %v", err)
	}

	return sp, nil
}

func (db *Database) SaveWhiteList(projc *ent.ServiceProvider, wlc *ent.WhiteListCreate) (*ent.ServiceProvider, error) {
	wl, err := wlc.Save(db.Ctx)

	if err != nil {
		return nil, err
	}

	proj, err := projc.
		Update().
		AddWhitelists(wl).
		Save(db.Ctx)

	return proj, err
}

func (db *Database) SaveAuthorizationCode(codeC *ent.AuthorizationCodeCreate) (*ent.AuthorizationCode, error) {
	code, err := codeC.Save(db.Ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to save authorization code: %v", err)
	}

	return code, err
}

func (db *Database) FindServiceProviderByHost(host string) (*ent.ServiceProvider, error) {
	return db.Client.ServiceProvider.
		Query().
		Where(serviceprovider.HostEQ(host)).
		Only(db.Ctx)
}
