package core

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	_ "sqlite3"

	"github.com/ochanoco/proxy/ent"
	"github.com/ochanoco/proxy/ent/hashchain"
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

	b, err := os.ReadFile(WHITELIST_PATH)
	if err != nil {
		return fmt.Errorf("failed to read white list (%v)\n=> %v", WHITELIST_PATH, err)
	}

	err = json.Unmarshal(b, &urls)

	if err != nil {
		return fmt.Errorf("failed to parse white list (%v)\n=> %v", WHITELIST_PATH, err)
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

func (db *Database) CreateHashChain(hash, signature []byte) *ent.HashChainCreate {
	chain := db.Client.HashChain.
		Create().
		SetHash(hash).
		SetSignature(signature)

	return chain
}

func (db *Database) SaveHashChain(l *ent.HashChainCreate) (*ent.HashChain, error) {
	code, err := l.Save(db.Ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to save save log: %v", err)
	}

	return code, err
}

func (db *Database) FindLastHashChain() (*ent.HashChain, error) {
	hash, err := db.Client.HashChain.
		Query().
		Order(ent.Desc(hashchain.FieldID)).
		Limit(1).
		All(db.Ctx)

	if err != nil {
		return nil, err
	}

	if len(hash) == 0 {
		return nil, nil
	}

	return hash[0], nil
}

func (db *Database) SelectAllServiceLogs() ([]*ent.ServiceLog, error) {
	return db.Client.ServiceLog.
		Query().
		All(db.Ctx)
}
