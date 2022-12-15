package database

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func (db *Database) MigrateWhiteList() error {
	var urls []string

	b, _ := os.ReadFile(WHITELIST_FILE)
	err := json.Unmarshal(b, &urls)

	if err != nil {
		log.Fatalf("failed to load migrate.json: %v", err)
		return err
	}

	projc := db.CreateServiceProvider(AUTH_PAGE_DOMAIN, AUTH_PAGE_DESTINATION)
	_, err = projc.Save(db.ctx)

	if err != nil {
		return fmt.Errorf("failed creating project: %v", err)
	}

	for _, url := range urls {
		wlc := db.CreateWhiteList(url)

		_, err := wlc.Save(db.ctx)

		if err != nil {
			return fmt.Errorf("failed add white list to project: %v", err)
		}
	}

	return nil
}
