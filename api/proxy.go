package main

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
)

func MainDirector(proxy *OchanocoProxy, req *http.Request) {
	loginRedirectURL, err := url.Parse(LOGIN_REDIRECT_PAGE_URL)
	if err != nil {
		GoToErrorPage("failed parse", err, req)
		return
	}

	project, err := proxy.Database.FindServiceProviderByHost(req.Host)

	if err != nil {
		msg := fmt.Sprintf("failed to get destination site (%s)", req.Host)
		GoToErrorPage(msg, err, req)
		return
	}

	req.URL.Scheme = "http"
	req.URL.Host = project.DestinationIP

	reg := regexp.MustCompile(`^\/ochanoco\/`)
	isOchanocoPath := reg.Match([]byte(req.URL.Path))

	isCleanContent := passIfCleanContent(req)
	isAuthed := authenticateRequest(req)

	if isOchanocoPath {
		req.URL.Scheme = "http"
		req.URL.Host = PROXYWEB_DOMAIN
		fmt.Println("match: hogehoge")

	} else if isCleanContent || isAuthed {
		req.Header.Set("User-Agent", "ochanoco")
		req.Header.Set("X-Ochanoco-Proxy-Token", "<proxy_token>")
	} else {
		fmt.Printf("cleanContent: %v, authed %v\n", isCleanContent, isAuthed)

		req.URL.Scheme = loginRedirectURL.Scheme
		req.URL.Host = loginRedirectURL.Host
		req.URL.Path = "/ochanoco/redirect"
	}

	LogReq(req)
}

func MainModifyResponse(proxy *OchanocoProxy, resp *http.Response) {
	fmt.Printf("=> %v\n", resp.Request.URL)
}
