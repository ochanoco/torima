package database

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func MigrateWhiteList() error {
	var urls []string

	b, _ := os.ReadFile(WHITELIST_FILE)
	err := json.Unmarshal(b, &urls)

	if err != nil {
		log.Fatalf("failed to load migrate.json: %v", err)
		return err
	}

	projc := CreateServiceProvider(DB, AUTH_PAGE_DOMAIN, AUTH_PAGE_DESTINATION)
	proj, nil := projc.Save(DB.ctx)

	if err != nil {
		return fmt.Errorf("failed creating project: %v", err)
	}

	for _, url := range urls {
		wl := CreateWhiteList(DB, url)
		proj, err = SaveWhiteListOnProj(DB, proj, wl)

		if err != nil {
			return fmt.Errorf("failed add white list to project: %v", err)
		}
	}

	return nil
}


