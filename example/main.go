package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/ochanoco/proxy/core"
	"github.com/ochanoco/proxy/serv"
)

func targetServ(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("X-Ochanoco-UserID")
	fmt.Fprintf(w,
		"<p>Hello! %v</p><br><a href='%v'>link</a>",
		userId,
		"/aaa?callback_path=/hello")
}

func main() {
	core.SCHEME = "http"
	core.CONFIG_FILE = "../config.yaml"
	core.STATIC_FOLDER = "../static"

	h := http.HandlerFunc(targetServ)
	server := httptest.NewServer(h)

	servUrl, _ := url.Parse(server.URL)
	proxyServ, err := serv.Run()

	if err != nil {
		panic(err)
	}

	proxyServ.Config.DefaultOrigin = servUrl.Host

	port := fmt.Sprintf(":%d", proxyServ.Config.Port)
	proxyServ.Engine.RunTLS(port, "./sample_keys/cert.pem", "./sample_keys/key.pem")
}
