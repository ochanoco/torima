package database

import (
	"encoding/json"
	"os"
	"testing"
)

func testMigrateWhiteList(t *testing.T) {
	var fwls []string

	db, err := setupForTest()
	if err != nil {
		t.Fatalf("failed to set up DB: %v", err)
	}

	defer db.client.Close()

	err = db.MigrateWhiteList()

	if err != nil {
		t.Fatalf("failed to migrate white list: %v", err)
		return
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
		return
	}

	for i := 0; i < len(dbwls); i++ {
		if fwls[i] != dbwls[i].Path {
			t.Fatalf("not match migration configuration and db data: %v", err)
		}
	}
}
