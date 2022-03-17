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

		wl := createWhiteList(db, "https://example.com")

		if err != nil {
			t.Errorf("failed creating white list: %v", err)
			return
		}

		projc := createProject(db, "<domain>", "<destination>", "<line_id>", "<name>")
		proj, nil := projc.Save(db.ctx)

		if err != nil {
			t.Errorf("failed creating project: %v", err)
			return
		}

		proj, err = savePageOnProj(db, proj, wl)

		if err != nil {
			t.Errorf("failed add white list to project: %v", err)
			return
		}

		log.Print(proj, wl)

		// log.Printf("project and white list created:\n    %v\n    %v", proj, wl)
	})
}
