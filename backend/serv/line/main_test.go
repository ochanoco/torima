package line

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/ochanoco/proxy/core"
)

func TestIntegration(t *testing.T) {
	core.DB_CONFIG = "file::memory:?cache=shared&_fk=1"
	core.SetupParsingUrl()

	proxyServ := Run()

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			fmt.Fprintf(w, "<p>404</p>")
			return
		}
		userId := r.Header.Get("X-Ochanoco-UserID")
		fmt.Fprintf(w, "<p>Hello! %v</p><br><a href='%v'>link</a>", userId, "/ochanoco/login?callback_path=/hello")
	})

	server := httptest.NewServer(h)
	defer server.Close()

	servUrl, _ := url.Parse(server.URL)

	sp := proxyServ.Database.Client.ServiceProvider.
		Create().
		SetHost("proxyweb:8080").
		SetDestinationIP(servUrl.Host)

	sp.SaveX(proxyServ.Database.Ctx)

	proxyServ.Engine.Run()
}
