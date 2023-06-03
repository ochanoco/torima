package core

import (
	"encoding/json"
	"log"
	"os"
	"testing"

	_ "sqlite3"
)

func TestDB(t *testing.T) {
	testMigrateWhiteList(t)
	testModel(t)
}

func testModel(t *testing.T) {
	db, err := InitDB(DB_CONFIG)

	if err != nil {
		t.Fatalf("model_test: %v", err)
	}

	defer db.Client.Close()

	wl := db.CreateWhiteList("https://example.com")

	if err != nil {
		t.Fatalf("failed creating white list: %v", err)
	}

	spc := db.CreateServiceProvider("<domain>", "<destination>")
	sp, err := db.SaveServiceProvider(spc)

	if err != nil {
		t.Fatalf("failed creating project: %v", err)
	}

	sp, err = db.SaveWhiteList(sp, wl)

	if err != nil {
		t.Fatalf("failed to add white list to project: %v", err)
	}

	log.Print(sp, wl)
}

func testMigrateWhiteList(t *testing.T) {
	var fwls []string

	db, err := InitDB(DB_CONFIG)
	if err != nil {
		t.Fatalf("failed to set up DB: %v", err)
	}

	defer db.Client.Close()

	err = db.MigrateWhiteList()

	if err != nil {
		t.Fatalf("failed to migrate white list:\n=> %v", err)
	}

	b, _ := os.ReadFile(WHITELIST_PATH)

	err = json.Unmarshal(b, &fwls)

	if err != nil {
		t.Fatalf("failed to load %v: %v", WHITELIST_PATH, err)
	}

	dbwls, err := db.Client.WhiteList.
		Query().All(db.Ctx)

	if err != nil {
		t.Fatalf("failed read white list: %v", err)
	}

	for i := 0; i < len(dbwls); i++ {
		if fwls[i] != dbwls[i].Path {
			t.Fatalf("not match migration configuration and db data: %v", err)
		}
	}
}
