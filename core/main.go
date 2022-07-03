package main

import (
	"net/http"
	"net/http/httputil"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	rp := httputil.ReverseProxy{
		Director:       director,
		ModifyResponse: modifyResponse,
	}

	serv := http.Server{
		Addr:    ":9000",
		Handler: &rp,
	}

	serv.ListenAndServe()
}
