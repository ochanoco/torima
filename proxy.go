package main

import (
	"log"
	"net/http"
	"net/url"
)

var LOGIN_PAGE_URL = "https://example.com"
var EXAMPLE_URL = "https://example.com"

func director(req *http.Request) {
	exampleURL, err := url.Parse(EXAMPLE_URL)
	if err != nil {
		log.Fatal(err)
	}

	loginURL, err := url.Parse(LOGIN_PAGE_URL)
	if err != nil {
		log.Fatal(err)
	}

	if authenticateRequest(req) {
		req.URL.Scheme = exampleURL.Scheme
		req.URL.Host = exampleURL.Host

		req.Header.Set("User-Agent", "bullet")
		req.Header.Set("X-BULLET-Proxy-Token", "<proxy_token>")
	} else {
		req.URL.Scheme = loginURL.Scheme
		req.URL.Host = loginURL.Host
	}
}

func modifyResponse(res *http.Response) error {
	// document, err := goquery.NewDocumentFromReader(res.Body)
	// if err != nil {
	// 	return err
	// }

	// body := document.Find("body")
	// body.SetHtml("hello")

	// html, err := document.Html()
	// if err != nil {
	// 	return err
	// }

	// b := []byte(html)
	// res.Body = ioutil.NopCloser(bytes.NewReader(b))
	// res.Header.Set("Content-Length", strconv.Itoa(len(b)))

	return nil
}
