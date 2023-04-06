package main

import (
	"fmt"
	"net/http"
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

	if ADD_USER_ID {
		req.Header.Set("X-Ochanoco-UserID", "1")
	}

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

func AuthDirector(proxy *OchanocoProxy, req *http.Request, c *gin.Context) bool {
	session := sessions.Default(c)
	userId := session.Get("user_id")

	switch userId.(type) {
	case int:
		req.URL.Scheme = ProxyRedirectUrl.Scheme
		req.URL.Host = ProxyRedirectUrl.Host
		req.URL.Path = "/ochanoco/login"

		return FINISHED
	default:
		return CONTINUE
	}
}
