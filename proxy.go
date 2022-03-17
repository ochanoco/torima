package main

import (
	"log"
	"net/http"
	"net/url"
)


func director(req *http.Request) {
	exampleURL, err := url.Parse(EXAMPLE_URL)
	if err != nil {
		log.Fatal(err)
	}
	req.URL.Scheme = exampleURL.Scheme
	req.URL.Host = exampleURL.Host

	loginRedirectURL, err := url.Parse(LOGIN_REDIRECT_PAGE_URL)
	if err != nil {
		log.Fatal(err)
	}

	isValid := passIfCleanContent(req)

	if authenticateRequest(req) {
		isValid = true
	}

	if isValid {
		req.Header.Set("User-Agent", "bullet")
		req.Header.Set("X-BULLET-Proxy-Token", "<proxy_token>")
	} else {
		req.URL.Scheme = loginRedirectURL.Scheme
		req.URL.Host = loginRedirectURL.Host
		req.URL.Path = "/redirect"
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
