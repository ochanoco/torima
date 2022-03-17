package main

import (
	"log"
	"testing"
)

func TestModel(t *testing.T) {
	DB_CONFIG = "file:ent?mode=memory&cache=shared&_fk=1"

	t.Run("test proxy", func(t *testing.T) {
		db, err := initDB()

		if err != nil {
			t.Errorf("model_test: %v", err)
		}

		page, err := createAndSavePage(db, "https://example.com", true)

		if err != nil {
			t.Errorf("failed creating page: %v", err)
			return
		}

		proj, err := createAndSaveProj(db, "<id>", "<name>")

		if err != nil {
			t.Errorf("failed creating project: %v", err)
			return
		}

		proj, err = addPageToProj(db, proj, page)

		if err != nil {
			t.Errorf("failed add page to project: %v", err)
			return
		}

		log.Printf("project and page created:\n    %v\n    %v", proj, page)
	})
}
