package main

import (
	"encoding/json"
	"log"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

const TEST_DB_PATH = "file:ent?mode=memory&cache=shared&_fk=1"

func TestDB(t *testing.T) {
	testMigrateWhiteList(t)
	testModel(t)
}

func testModel(t *testing.T) {
	db, err := InitDB(TEST_DB_PATH)

	if err != nil {
		t.Fatalf("model_test: %v", err)
	}

	defer db.client.Close()

	wl := db.CreateWhiteList("https://example.com")

	if err != nil {
		t.Fatalf("failed creating white list: %v", err)
	}

	projc := db.CreateServiceProvider("<domain>", "<destination>")
	proj, err := projc.Save(db.ctx)

	if err != nil {
		t.Fatalf("failed creating project: %v", err)
	}

	proj, err = db.SaveWhiteList(proj, wl)

	if err != nil {
		t.Fatalf("failed add white list to project: %v", err)
	}

	log.Print(proj, wl)
}

func testMigrateWhiteList(t *testing.T) {
	var fwls []string

	db, err := InitDB(TEST_DB_PATH)
	if err != nil {
		t.Fatalf("failed to set up DB: %v", err)
	}

	defer db.client.Close()

	err = db.MigrateWhiteList()

	if err != nil {
		t.Fatalf("failed to migrate white list: %v", err)
	}

	b, _ := os.ReadFile(WHITELIST_FILE)
	err = json.Unmarshal(b, &fwls)

	if err != nil {
		t.Fatalf("failed to load migrate.json: %v", err)
	}

	dbwls, err := db.client.WhiteList.
		Query().All(db.ctx)

	if err != nil {
		t.Fatalf("failed read white list: %v", err)
	}

	for i := 0; i < len(dbwls); i++ {
		if fwls[i] != dbwls[i].Path {
			t.Fatalf("not match migration configuration and db data: %v", err)
		}
	}
}
