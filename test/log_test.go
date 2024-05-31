package test

import (
	"fmt"
	"testing"

	"github.com/ochanoco/torima/core"
	"github.com/ochanoco/torima/extension/directors"
	"github.com/ochanoco/torima/test/tools"
)

func TestExtensionLog(t *testing.T) {
	core.DB_TYPE = "sqlite3"
	core.DB_CONFIG = "../data/test.db?_fk=1"
	core.SECRET = "test_secret"

	logger := tools.ExtensionLogger{}

	directors := core.TorimaDirectors{
		directors.DefaultRouteDirector,
		directors.DefaultRouteDirector,
	}

	directors = logger.InjectDirectors(directors)

	if len(directors) != 5 {
		t.Errorf("InjectDirector failed")
	}

	fmt.Printf("Directors: %v\n", directors)

	c, _ := directorSample(t)
	c.Proxy.Directors = directors

	c.Proxy.Director(c.Target, c.GinContext)
}
