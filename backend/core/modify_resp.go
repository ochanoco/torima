package core

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func MainModifyResponse(proxy *OchanocoProxy, resp *http.Response) {
	fmt.Printf("=> %v\n", resp.Request.URL)
}

func ResponseLogModifyResponse(proxy *OchanocoProxy, res *http.Response, c *gin.Context) bool {
	log.Printf("LogDirector: start")
	// Current date in format "2006-01-02T15:04:05Z07:00"
	tim := time.Now()

	db := proxy.Database
	if db == nil {
		log.Printf("LogDirector: db is nil")
		return false
	}

	l := db.Client.ServiceLog.Create()
	l.SetTime(tim)

	headerJson, err := DumpHeader(res.Header)
	if err != nil {
		log.Printf("ResponseLogDirector: failed to dump headers to json")
		return false
	}

	log.Printf("ResponseLogDirector: ========== start header ==========")
	log.Print(headerJson)
	log.Printf("ResponseLogDirector: ==========  end header  ==========")
	l.SetHeaders(headerJson)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("ResponseLogDirector: non-nil error while reading request body: %v", err)
		return false
	}
	res.Body.Close()
	res.Body = io.NopCloser(bytes.NewBuffer(body))

	l.SetBody(body)

	saved, err := l.Save(context.TODO())
	if err != nil {
		log.Printf("ResponseLogDirector: failed to save: %v", err)
		return false
	}

	log.Printf("ResponseLogDirector: log saved:  %v", saved)

	log.Printf("ResponseLogDirector: end")
	// hello

	return true
}
