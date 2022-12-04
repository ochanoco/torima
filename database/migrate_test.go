package database

import (
	"encoding/json"
	"log"
	"os"
	"testing"
)

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
