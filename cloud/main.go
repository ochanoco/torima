package main

import (
	"log"
	"net/http"
	"net/http/httputil"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	_, err := initDB()

	if err != nil {
		log.Panicf("failed init db: %v", err)
	}

	migrateWhiteList()

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
