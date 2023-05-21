package core

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

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

	db := proxy.Database
	if db == nil {
		log.Printf("LogDirector: db is nil")
		return false
	}

	l := db.Client.ServiceLog.Create()

	headerJson, err := DumpHeader(req.Header)
	if err != nil {
		log.Printf("LogDirector: failed to dump headers to json")
		return false
	}

	log.Printf("LogDirector: ========== start header ==========")
	log.Print(headerJson)
	log.Printf("LogDirector: ==========  end header  ==========")
	l.SetHeaders(headerJson)

	if req.Method == http.MethodGet || req.Method == http.MethodHead || req.Method == http.MethodOptions || req.Body == nil {
		log.Printf("LogDirector: no-body method")
		l.SetBody(nil)
	} else {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			log.Printf("LogDirector: non-nil error while reading request body: %v", err)
			return false
		}
		req.Body.Close()
		req.Body = io.NopCloser(bytes.NewBuffer(body))

		l.SetBody(body)
	}

	saved, err := l.Save(req.Context())
	if err != nil {
		log.Printf("LogDirector: failed to save: %v", err)
		return false
	}

	log.Printf("LogDirector: log saved:  %v", saved)

	log.Printf("LogDirector: end")
	// hello

	return true
}
