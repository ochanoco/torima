package database

import (
	"log"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func testModel(t *testing.T) {
	db, err := setupForTest()

	if err != nil {
		t.Errorf("model_test: %v", err)
	}

	defer db.client.Close()

	wl := db.CreateWhiteList("https://example.com")

	if err != nil {
		t.Errorf("failed creating white list: %v", err)
		return
	}

	projc := db.CreateServiceProvider("<domain>", "<destination>")
	proj, nil := projc.Save(db.ctx)

	if err != nil {
		t.Errorf("failed creating project: %v", err)
		return
	}

	proj, err = db.SaveWhiteList(proj, wl)

	if err != nil {
		t.Errorf("failed add white list to project: %v", err)
		return
	}

	log.Print(proj, wl)
}
