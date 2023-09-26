package core

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

func RouteDirector(host string, proxy *OchanocoProxy, req *http.Request, c *gin.Context) (bool, error) {
	req.URL.Host = host

	req.Header.Set("User-Agent", "ochanoco")
	req.Header.Set("X-Ochanoco-Proxy-Token", SECRET)

	req.URL.Scheme = proxy.Config.Scheme

	return CONTINUE, nil
}

func DefaultRouteDirector(proxy *OchanocoProxy, req *http.Request, c *gin.Context) (bool, error) {
	if strings.HasPrefix(req.URL.Path, "/ochanoco/") {
		return CONTINUE, nil
	}

	host := proxy.Config.DefaultOrigin

	if host == "" {
		err := fmt.Errorf("failed to get destination config (%s)", host)
		return FINISHED, err
	}

	return RouteDirector(host, proxy, req, c)
}

func ThirdPartyDirector(proxy *OchanocoProxy, req *http.Request, c *gin.Context) (bool, error) {
	path := strings.Split(req.URL.Path, "/")
	hasRedirectPrefix := strings.HasPrefix(req.URL.Path, "/ochanoco/redirect/")

	if !hasRedirectPrefix || len(path) < 3 {
		return CONTINUE, nil
	}

	for _, origin := range proxy.Config.ProtectionScope {
		if origin == path[3] {
			req.Host = origin
			req.URL.Host = origin

			p := strings.Join(path[4:], "/")
			req.URL.Path = "/" + p

			req.URL.Scheme = "https"
			return RouteDirector(origin, proxy, req, c)
		}
	}

	return CONTINUE, nil
}

func AuthDirector(proxy *OchanocoProxy, req *http.Request, c *gin.Context) (bool, error) {
	session := sessions.Default(c)
	userId := session.Get("userId")

	if userId != nil {
		req.Header.Set("X-Ochanoco-UserID", userId.(string))
		return CONTINUE, nil
	}

	if req.Method == "GET" {
		if req.URL.Path == "/" {
			return CONTINUE, nil
		}

		if slices.Contains(proxy.Config.WhiteListPath, req.URL.Path) {
			return CONTINUE, nil
		}
	}

	return FINISHED, makeError(fmt.Errorf(""), unauthorizedErrorTag)
}

func LogDirector(proxy *OchanocoProxy, req *http.Request, c *gin.Context) (bool, error) {
	request, err := httputil.DumpRequest(req, true)

	if err != nil {
		err = makeError(err, "failed to dump headers to json: %v")
		return FINISHED, err
	}

	splited := bytes.Split(request, []byte("\r\n\r\n"))

	header := splited[0]
	headerLen := len(header)

	body := request[headerLen:]

	l := proxy.Database.CreateRequestLog(string(header), body)
	_, err = l.Save(proxy.Database.Ctx)

	if err != nil {
		err = makeError(err, "failed to save request: %v")
		return FINISHED, err
	}

	return CONTINUE, err
}
