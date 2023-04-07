package example

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/ochanoco/proxy/core"
)

func targetServ(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("X-Ochanoco-UserID")
	fmt.Fprintf(w, "<p>Hello! %v</p><br><a href='%v'>link</a>", userId, "/ochanoco/login?callback_path=/hello")
}

func prepare() *core.OchanocoProxy {
	core.DB_CONFIG = "file::memory:?cache=shared&_fk=1"
	secret := "testest"

	core.SetupParsingUrl()

	h := http.HandlerFunc(targetServ)
	server := httptest.NewServer(h)

	servUrl, _ := url.Parse(server.URL)
	proxyServ := core.ProxyServer(secret)

	sp := proxyServ.Database.Client.ServiceProvider.
		Create().
		SetHost("127.0.0.1:8080").
		SetDestinationIP(servUrl.Host)

	sp.SaveX(proxyServ.Database.Ctx)

	return proxyServ
}
