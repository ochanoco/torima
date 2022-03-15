package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

var LOGIN_PAGE_URL = "https://example.com"

func director(req *http.Request) {
	rpURL, err := url.Parse(LOGIN_PAGE_URL)

	if err != nil {
		log.Fatal(err)
	}

	req.URL.Scheme = rpURL.Scheme
	req.URL.Host = rpURL.Host

	req.Header.Set("User-Agent", "tuaset")
	req.Header.Set("X-TUASET-User-ID", "<user_id>")
	req.Header.Set("X-TUASET-Proxy-Token", "<token>")
}

func modifyResponse(res *http.Response) error {
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err
	}

	body := document.Find("body")
	body.SetHtml("hello")

	html, err := document.Html()
	if err != nil {
		return err
	}

	b := []byte(html)
	res.Body = ioutil.NopCloser(bytes.NewReader(b))
	res.Header.Set("Content-Length", strconv.Itoa(len(b)))

	return nil
}
