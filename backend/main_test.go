package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestIntegration(t *testing.T) {
	DB_CONFIG = "file::memory:?cache=shared&_fk=1"
	secret := "testest"
	setupParsingUrl()

	proxyServ := ProxyServer(secret)
	authServ := AuthServer(secret, proxyServ)

	if os.Getenv("TEST_INTEGRATION") != "1" {
		t.Skip("Skipping testing in All test")
	}

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId := r.Header.Get("X-Ochanoco-UserID")
		fmt.Fprintf(w, "<p>Hello! %v</p><br><a href='%v'>link</a>", userId, "/ochanoco/redirect?callback_path=/hello")
	})

	server := httptest.NewServer(h)
	defer server.Close()

	servUrl := parseURL(t, server.URL)

	sp := proxyServ.Database.client.ServiceProvider.
		Create().
		SetHost("127.0.0.1:8080").
		SetDestinationIP(servUrl.Host)

	sp.SaveX(proxyServ.Database.ctx)

	go authServ.Run(AUTH_PORT)
	proxyServ.Engine.Run(PROXY_PORT)
}
