package database

import (
	"log"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func setupForTest() {
	DB_CONFIG = "file:ent?mode=memory&cache=shared&_fk=1"
	dbl, err := InitDB()

	if err != nil {
		log.Panicf("failed init db: %v", err)
	}

	DB = dbl
}

func TestModel(t *testing.T) {
	setupForTest()

	t.Run("test model", func(t *testing.T) {
		db, err := InitDB()

		if err != nil {
			t.Errorf("model_test: %v", err)
		}

		wl := CreateWhiteList(db, "https://example.com")

		if err != nil {
			t.Errorf("failed creating white list: %v", err)
			return
		}

		projc := CreateServiceProvider(db, "<domain>", "<destination>")
		proj, nil := projc.Save(db.ctx)

		if err != nil {
			t.Errorf("failed creating project: %v", err)
			return
		}

		proj, err = SaveWhiteListOnProj(db, proj, wl)

		if err != nil {
			t.Errorf("failed add white list to project: %v", err)
			return
		}

		log.Print(proj, wl)
	})
}
