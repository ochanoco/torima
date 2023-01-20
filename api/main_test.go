package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

func TestIntegration(t *testing.T) {
	if os.Getenv("TEST_INTEGRATION") != "1" {
		t.Skip("Skipping testing in All test")
	}

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<a href='%v'>link</a>", "/ochanoco/redirect")
	})

	server := httptest.NewServer(h)
	defer server.Close()

	db, err := InitDB(DB_CONFIG)
	if err != nil {
		log.Fatalf("failed to init db: %v", err)
	}

	servUrl, err := url.Parse(server.URL)
	if err != nil {
		log.Fatal(err)
	}

	sp := db.client.ServiceProvider.
		Create().
		SetHost("127.0.0.1:9000").
		SetDestinationIP(servUrl.Host)

	sp.SaveX(db.ctx)

	main()
}
