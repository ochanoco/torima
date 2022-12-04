package database

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/ochanoco/database/ent"
)

func SaveWhiteListOnProj(db *Database, projc *ent.ServiceProvider, wlc *ent.WhiteListCreate) (*ent.ServiceProvider, error) {
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
