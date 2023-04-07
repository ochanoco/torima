package example

import (
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/ochanoco/proxy/auth/line"
	"github.com/ochanoco/proxy/core"
)

func Main() {
	core.DB_CONFIG = "file::memory:?cache=shared&_fk=1"

	core.SetupParsingUrl()

	h := http.HandlerFunc(targetServ)
	server := httptest.NewServer(h)

	servUrl, _ := url.Parse(server.URL)
	proxyServ := line.Run()

	sp := proxyServ.Database.Client.ServiceProvider.
		Create().
		SetHost("127.0.0.1:8080").
		SetDestinationIP(servUrl.Host)

	sp.SaveX(proxyServ.Database.Ctx)

	proxyServ.Engine.Run(core.PROXY_PORT)
}

const LINE_NAME = "line_example"
