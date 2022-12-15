package database

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/ochanoco/database/ent"
)

func (db *Database) CreateWhiteList(path string) *ent.WhiteListCreate {
	wl := db.client.WhiteList.
		Create().
		SetPath(path)

	return wl
}

func (db *Database) CreateServiceProvider(host string, destinationIP string) *ent.ServiceProviderCreate {
	proj := db.client.ServiceProvider.
		Create().
		SetHost(host).
		SetDestinationIP(destinationIP)

	return proj
}
