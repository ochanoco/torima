package main

import (
	"log"
	"net/http"
	"net/url"

	"github.com/tracer-silver-bullet/tracer-silver-bullet/proxy/ent/project"
)

func goToErrorPage(msg string, err error, req *http.Request) {
	log.Fatalf(msg, err)

	errorPageURL, err := url.Parse(ERROR_PAGE_URL)

	if err != nil {
		log.Fatalf("parse error: %v", err)
	}

	req.URL.Scheme = errorPageURL.Scheme
	req.URL.Host = errorPageURL.Host
	req.URL.Path = "/404"
}

func director(req *http.Request) {
	loginRedirectURL, err := url.Parse(LOGIN_REDIRECT_PAGE_URL)

	if err != nil {
		goToErrorPage("failed parse", err, req)
		return
	}

	project, err := db.client.Project.
		Query().
		Where(project.DomainEQ(req.Host)).
		Only(db.ctx)

	if err != nil {
		goToErrorPage("failed find site", err, req)
		return
	}

	req.URL.Scheme = "http"
	req.URL.Host = project.Destination

	isValid := passIfCleanContent(req)
	isValid = isValid || authenticateRequest(req)

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
