package proxy

import (
	"fmt"
	"net/http"
	"net/url"
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

	LogReq(req)
}
