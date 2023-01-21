package main

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
)

const CONTINUE = true
const FINISHED = false

var DEFAULT_DIRECTORS = []OchanocoDirector{
	SpecialPathDirector,
	AuthDirector,
	DefaultDirector,
}

var DEFAULT_MODIFY_RESPONSES = []OchanocoModifyResponse{}

/**
 * Directors is a list of functions that modify the
 * request before it is sent to the target server.
 **/
func (proxy *OchanocoProxy) Director(req *http.Request) {
	for _, d := range proxy.Directors {
		d(proxy, req)
	}
}

/**
  * ModifyResponses is a list of functions that modify the
  * response before it is sent to the client.
**/
func (proxy *OchanocoProxy) ModifyResponse(res *http.Response) error {
	for _, mR := range proxy.ModifyResponses {
		mR(proxy, res)
	}

	return nil
}

func SpecialPathDirector(proxy *OchanocoProxy, req *http.Request) bool {
	reg := regexp.MustCompile(`^\/ochanoco\/`)
	isOchanocoPath := reg.Match([]byte(req.URL.Path))

	if isOchanocoPath {
		req.URL.Scheme = "http"
		req.URL.Host = PROXYWEB_DOMAIN

		return FINISHED
	}

	return CONTINUE
}

func AuthDirector(proxy *OchanocoProxy, req *http.Request) bool {
	loginRedirectURL, err := url.Parse(LOGIN_REDIRECT_PAGE_URL)

	if err != nil {
		GoToErrorPage("failed parse", err, req)
		return FINISHED
	}

	isCleanContent := passIfCleanContent(req)
	isAuthed := authenticateRequest(req)

	fmt.Printf("cleanContent: %v, authed %v\n", isCleanContent, isAuthed)

	if !isCleanContent && !isAuthed {
		req.URL.Scheme = loginRedirectURL.Scheme
		req.URL.Host = loginRedirectURL.Host
		req.URL.Path = "/ochanoco/redirect"

		return FINISHED
	}

	return CONTINUE
}

func DefaultDirector(proxy *OchanocoProxy, req *http.Request) bool {
	project, err := proxy.Database.FindServiceProviderByHost(req.Host)

	if err != nil {
		msg := fmt.Sprintf("failed to get destination site (%s)", req.Host)
		GoToErrorPage(msg, err, req)
		return FINISHED
	}

	req.URL.Scheme = "http"
	req.URL.Host = project.DestinationIP

	req.Header.Set("User-Agent", "ochanoco")
	req.Header.Set("X-Ochanoco-Proxy-Token", "<proxy_token>")

	return CONTINUE
}

func MainModifyResponse(proxy *OchanocoProxy, resp *http.Response) {
	fmt.Printf("=> %v\n", resp.Request.URL)
}
