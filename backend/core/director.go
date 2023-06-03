package core

import (
	"fmt"
	"net/http"
	"os"
	"regexp"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var cleanContentPattern = regexp.MustCompile(`.+\.(html|css|js|jpg|png|gif)`)

func RouteDirector(host string, proxy *OchanocoProxy, req *http.Request, c *gin.Context) bool {
	req.URL.Host = host

	req.Header.Set("User-Agent", "ochanoco")
	req.Header.Set("X-Ochanoco-Proxy-Token", "<proxy_token>")

	session := sessions.Default(c)
	userId := session.Get("userId")

	fmt.Printf("data: %v\n", userId)

	if ADD_USER_ID {
		switch userId.(type) {
		case string:
			req.Header.Set("X-Ochanoco-UserID", userId.(string))

		default:
			req.Header.Set("X-Ochanoco-UserID", "nil")
		}
	}

	return CONTINUE
}

func EnvRouteDirector(proxy *OchanocoProxy, req *http.Request, c *gin.Context) bool {
	host := os.Getenv("OCHANOCO_DESTINATION")

	if host == "" {
		msg := fmt.Errorf("failed to get destination site on environment variable named 'DESTINATION' (%s)", req.Host)
		GoToErrorPage("", msg, req)
		return FINISHED
	}

	return RouteDirector(host, proxy, req, c)
}

func CloudRouteDirector(proxy *OchanocoProxy, req *http.Request, c *gin.Context) bool {
	project, err := proxy.Database.FindServiceProviderByHost(req.Host)

	if err != nil {
		msg := fmt.Sprintf("failed to get destination site (%s)", req.Host)
		GoToErrorPage(msg, err, req)
		return FINISHED
	}

	return RouteDirector(project.DestinationIP, proxy, req, c)
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
	case string:
		req.URL.Scheme = ProxyRedirectUrl.Scheme
		req.URL.Host = ProxyRedirectUrl.Host
		req.URL.Path = "/ochanoco/login"

		return FINISHED
	default:
		return CONTINUE
	}
}

func LogDirector(proxy *OchanocoProxy, req *http.Request, c *gin.Context) bool {
	_, err := logToDB(req.Header, req.Body, proxy, c)

	if err != nil {
		fmt.Printf("LogModifyResponse: %v\n", err)
		return false
	}
	return true
}
