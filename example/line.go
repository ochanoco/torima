package example

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/ochanoco/proxy/serv"
)

func Main() {
	h := http.HandlerFunc(targetServ)
	server := httptest.NewServer(h)

	servUrl, _ := url.Parse(server.URL)
	proxyServ, err := serv.Run()

	if err != nil {
		panic(err)
	}

	proxyServ.Config.DefaultOrigin = servUrl.Host

	port := fmt.Sprintf(":%d", proxyServ.Config.Port)
	proxyServ.Engine.Run(port)
}

const LINE_NAME = "line_example"
