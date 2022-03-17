package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	_ "github.com/mattn/go-sqlite3"
	"github.com/tracer-silver-bullet/tracer-silver-bullet/proxy/ent"
)

func CreatePage(ctx context.Context, client *ent.Client) (*ent.Page, error) {
	u, err := client.Page.
		Create().
		SetURL("https://google.com").
		SetSkip(false).
		SetProjectID(1).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	ctx := context.Background()

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	CreatePage(ctx, client)

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
