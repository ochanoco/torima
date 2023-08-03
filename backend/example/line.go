package example

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/ochanoco/proxy/serv/line"
)

func Main() {
	// core.DB_CONFIG = "file::memory:?cache=shared&_fk=1"
	h := http.HandlerFunc(targetServ)
	server := httptest.NewServer(h)

	servUrl, _ := url.Parse(server.URL)
	proxyServ, err := line.Run()

	if err != nil {
		panic(err)
	}

	proxyServ.Config.DefaultOrigin = servUrl.Host

	port := fmt.Sprintf(":%d", proxyServ.Config.Port)
	proxyServ.Engine.Run(port)
}

const LINE_NAME = "line_example"
