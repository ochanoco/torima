package database

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/ochanoco/database/ent"
)

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
