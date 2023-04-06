package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestIntegration(t *testing.T) {
	if os.Getenv("TEST_INTEGRATION") != "1" {
		t.Skip("Skipping testing in All test")
	}

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId := r.Header.Get("X-Ochanoco-UserID")
		fmt.Fprintf(w, "<p>Hello! %v</p><br><a href='%v'>link</a>", userId, "/ochanoco/redirect?callback_path=/hello")
	})

	server := httptest.NewServer(h)
	defer server.Close()

	db, err := InitDB(DB_CONFIG)
	if err != nil {
		log.Fatalf("failed to init db: %v", err)
	}

	servUrl := parseURL(t, server.URL)

	sp := db.client.ServiceProvider.
		Create().
		SetHost("127.0.0.1:8080").
		SetDestinationIP(servUrl.Host)

	sp.SaveX(db.ctx)

	main()
}
