package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/ochanoco/ochano.co-auth/proxy/ent/project"
)

func logReq(req *http.Request) {
	fmt.Printf("[%s] %s%s\n=> %s%s\n\n", req.Method, req.Host, req.RequestURI, req.URL.Host, req.URL.Path)
}

func goToErrorPage(msg string, err error, req *http.Request) {
	fmt.Fprintln(os.Stderr, msg, err)

	errorPageURL, err := url.Parse(ERROR_PAGE_URL)

	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}

	req.URL.Scheme = errorPageURL.Scheme
	req.URL.Host = errorPageURL.Host
	req.URL.Path = "/404?msg=" + msg

	logReq(req)
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
		msg := fmt.Sprintf("failed to get destination site (%s)", req.Host)
		goToErrorPage(msg, err, req)
		return
	}

	req.URL.Scheme = "http"
	req.URL.Host = project.Destination

	isCleanContent := passIfCleanContent(req)
	isAuthed := authenticateRequest(req)

	if isCleanContent || isAuthed {
		req.Header.Set("User-Agent", "ochanoco")
		req.Header.Set("X-Ochanoco-Proxy-Token", "<proxy_token>")
	} else {
		req.URL.Scheme = loginRedirectURL.Scheme
		req.URL.Host = loginRedirectURL.Host
		req.URL.Path = fmt.Sprintf("/redirect?clean=%v&authed=%v", isCleanContent, isAuthed)
	}

	logReq(req)
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
