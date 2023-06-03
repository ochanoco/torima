package example

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"

	"github.com/ochanoco/proxy/core"
	"github.com/ochanoco/proxy/serv/line"
)

func Main() {
	// core.DB_CONFIG = "file::memory:?cache=shared&_fk=1"

	core.SetupParsingUrl()

	h := http.HandlerFunc(targetServ)
	server := httptest.NewServer(h)

	servUrl, _ := url.Parse(server.URL)
	proxyServ := line.Run()

	os.Setenv("OCHANOCO_DESTINATION", servUrl.Host)

	proxyServ.Engine.Run(core.PROXY_PORT)
}

const LINE_NAME = "line_example"
