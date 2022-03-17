package main

import (
	"log"
	"testing"
)

func TestModel(t *testing.T) {
	// DB_CONFIG = "file:ent?mode=memory&cache=shared&_fk=1"

	t.Run("test proxy", func(t *testing.T) {
		db, err := init_db()

		if err != nil {
			t.Errorf("model_test: %v", err)
		}

		page, err := db.client.Page.
			Create().
			SetURL("https://example.com").
			SetSkip(false).
			Save(db.ctx)

		if err != nil {
			t.Errorf("failed creating page: %v", err)
			return
		}

		proj, err := db.client.Project.
			Create().
			SetLineID("1").
			SetName("taro").
			AddPages(page).
			Save(db.ctx)

		if err != nil {
			t.Errorf("failed creating project: %v", err)
			return
		}

		log.Printf("project and page created:\n    %v\n    %v", proj, page)
	})
}
