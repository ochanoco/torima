package database

import (
	"encoding/json"
	"log"
	"os"
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

func TestMigrateWhiteList(t *testing.T) {
	var fwls []string

	setupForTest()

	err := MigrateWhiteList()

	if err != nil {
		t.Errorf("failed to migrate white list: %v", err)
		return
	}

	b, _ := os.ReadFile(WHITELIST_FILE)
	err = json.Unmarshal(b, &fwls)

	if err != nil {
		log.Fatalf("failed to load migrate.json: %v", err)
	}

	t.Run("test model", func(t *testing.T) {
		dbwls, err := DB.client.WhiteList.
			Query().All(DB.ctx)

		if err != nil {
			t.Errorf("failed read white list: %v", err)
			return
		}

		for i := 0; i < len(dbwls); i++ {
			if fwls[i] != dbwls[i].Path {
				t.Errorf("not match migration configuration and db data: %v", err)

			}
		}
	})

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
