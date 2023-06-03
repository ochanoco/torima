package core

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"

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
	log.Printf("LogDirector: start")
	// Current date in format "2006-01-02T15:04:05Z07:00"
	tim := time.Now()

	db := proxy.Database
	if db == nil {
		log.Printf("RequestLogDirector: db is nil")
		return false
	}

	l := db.Client.ServiceLog.Create()

	l.SetTime(tim)

	headerJson, err := DumpHeader(req.Header)
	if err != nil {
		log.Printf("RequestLogDirector: failed to dump headers to json")
		return false
	}

	log.Printf("RequestLogDirector: ========== start header ==========")
	log.Print(headerJson)
	log.Printf("RequestLogDirector: ==========  end header  ==========")
	l.SetHeaders(headerJson)

	// There are kinds of methods which does not have bodies (i.e., GET, HEAD, OPTIONS, TRACE).
	if req.Body == nil {
		log.Printf("RequestLogDirector: no-body method")
		l.SetBody(nil)
	} else {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			log.Printf("RequestLogDirector: non-nil error while reading request body: %v", err)
			return false
		}
		req.Body.Close()
		req.Body = io.NopCloser(bytes.NewBuffer(body))

		l.SetBody(body)
	}

	saved, err := l.Save(req.Context())
	if err != nil {
		log.Printf("RequestLogDirector: failed to save: %v", err)
		return false
	}

	log.Printf("RequestLogDirector: log saved:  %v", saved)

	log.Printf("RequestLogDirector: end")
	// hello

	return true
}
