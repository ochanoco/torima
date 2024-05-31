package core

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func MainModifyResponse(proxy *TorimaProxy, res *http.Response) {
	fmt.Printf("=> %v\n", res.Request.URL)
}

func InjectHTMLModifyResponse(html string, c *TorimaModifyResponsePackageContext) (TorimaPackageStatus, error) {
	document, err := goquery.NewDocumentFromReader(c.Target.Body)
	if err != nil {
		log.Fatal(err)
	}

	document.Find("body").AppendHtml(html)

	html, err = document.Html()
	if err != nil {
		return ForceStop, err
	}

	// fmt.Printf("%v", html)

	b := []byte(html)
	c.Target.Body = io.NopCloser(bytes.NewReader(b))

	c.Target.Header.Set("Content-Length", strconv.Itoa(len(b)))
	c.Target.ContentLength = int64(len(b))

	return Keep, nil
}

func InjectServiceWorkerModifyResponse(c *TorimaModifyResponsePackageContext) (TorimaPackageStatus, error) {
	contentType := c.Target.Header.Get("Content-Type")

	if contentType != "text/html; charset=utf-8" {
		return Keep, nil
	}

	html := scripts + "\n"

	return InjectHTMLModifyResponse(html, c)
}
