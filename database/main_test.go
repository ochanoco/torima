package database

import (
	"fmt"
	"testing"
)

func setupForTest() (*Database, error) {
	DB_CONFIG = "file:ent?mode=memory&cache=shared&_fk=1"

	dbl, err := InitDB(DB_CONFIG)

	if err != nil {
		return nil, fmt.Errorf("failed init db: %v", err)
	}

	return dbl, nil
}

func TestMain(t *testing.T) {
	t.Run("test model", testModel)
	t.Run("test migrate whiteTestMain", testMigrateWhiteList)
}
