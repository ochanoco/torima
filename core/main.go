package main

import (
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	directors := []OchanocoDirector{}
	modifyResponses := []OchanocoResponse{}

	db, err := InitDB("./sqlite3.db")
	if err != nil {
		log.Fatalf("failed to init db: %v", err)
	}

	proxy := NewOchancoProxy(directors, modifyResponses, db)
	server := http.Server{
		Addr:    ":9000",
		Handler: proxy.ReverseProxy,
	}

	server.ListenAndServe()
}
