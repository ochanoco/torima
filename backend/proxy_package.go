package main

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var patternSpecialPath = regexp.MustCompile(`^\/ochanoco\/`)

var cleanContentPattern = regexp.MustCompile(`.+\.(html|css|js|jpg|png|gif)`)

func MainDirector(proxy *OchanocoProxy, req *http.Request, c *gin.Context) bool {
	project, err := proxy.Database.FindServiceProviderByHost(req.Host)

	if err != nil {
		msg := fmt.Sprintf("failed to get destination site (%s)", req.Host)
		GoToErrorPage(msg, err, req)
		return FINISHED
	}

	req.URL.Host = project.DestinationIP

	req.Header.Set("User-Agent", "ochanoco")
	req.Header.Set("X-Ochanoco-Proxy-Token", "<proxy_token>")

	return CONTINUE
}

func CleanContentDirector(proxy *OchanocoProxy, req *http.Request, c *gin.Context) bool {
	if req.Method != "GET" {
		// If the request is not GET, the request is passed.
		return CONTINUE
	}

	if req.RequestURI == "/" || cleanContentPattern.MatchString(req.URL.Path) {
		// If the request is for static content, the request is passed.
		req.URL.Path = cleanContentPattern.FindString(req.URL.Path)
		return FINISHED
	}

	return CONTINUE
}

func LoginPathDirector(proxy *OchanocoProxy, req *http.Request, c *gin.Context) bool {
	fmt.Printf("Path: %v\n", req.URL.Path)

	if req.URL.Path != "/ochanoco/redirect" {
		return CONTINUE
	}

	req.URL.Host = PROXYWEB_HOST

	callback_path := req.URL.Query().Get("callback_path")

	if callback_path == "" {
		callback_path = "/"
	}

	session := sessions.Default(c)
	session.Set("callback_path", callback_path)

	return FINISHED
}

func CallbackPathDirector(proxy *OchanocoProxy, req *http.Request, c *gin.Context) bool {
	if req.URL.Path != "/ochanoco/callback" {
		return CONTINUE
	}

	token := req.URL.Query().Get("authorization_code")
	if token == "" {
		panic("failed to get authorization_code")
	}

	session := sessions.Default(c)
	session.Set("authorization_code", token)

	if session.Save() != nil {
		panic("failed to save authorization_code")
	}

	callbackPath := session.Get("callback_path")
	if callbackPath == nil {
		callbackPath = "/"
	}

	req.URL.Path = callbackPath.(string)
	req.URL.Host = req.Host
	req.URL.RawQuery = ""

	return FINISHED
}

func ProxyPageDirector(proxy *OchanocoProxy, req *http.Request, c *gin.Context) bool {
	isOchanocoPath := patternSpecialPath.Match([]byte(req.URL.Path))

	if isOchanocoPath {
		req.URL.Host = PROXYWEB_HOST

		return FINISHED
	}

	return CONTINUE
}

func AuthDirector(proxy *OchanocoProxy, req *http.Request, c *gin.Context) bool {
	loginRedirectURL, err := url.Parse(PROXY_REDIRECT_URL)

	if err != nil {
		GoToErrorPage("failed parse", err, req)
		return FINISHED
	}

	isAuthed := authenticateRequest(req)

	if isAuthed {
		req.URL.Scheme = loginRedirectURL.Scheme
		req.URL.Host = loginRedirectURL.Host
		req.URL.Path = "/ochanoco/redirect"

		return FINISHED
	}

	return CONTINUE
}

func MainModifyResponse(proxy *OchanocoProxy, resp *http.Response) {
	fmt.Printf("=> %v\n", resp.Request.URL)
}
