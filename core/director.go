package core

import (
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
	req.Header.Set("X-Ochanoco-Proxy-Token", "<proxy_token>")

	// req.URL.Scheme = "https"

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

	for _, origin := range proxy.Config.AcceptedOrigins {
		if origin == path[3] {
			req.Host = origin
			req.URL.Host = origin

			p := strings.Join(path[4:], "/")
			req.URL.Path = "/" + p

			return RouteDirector(origin, proxy, req, c)
		}
	}

	return CONTINUE, nil
}

func AuthDirector(proxy *OchanocoProxy, req *http.Request, c *gin.Context) (bool, error) {
	session := sessions.Default(c)
	userId := session.Get("userId")

	if userId != nil {
		return CONTINUE, nil
	}

	if req.Method == "GET" {
		if req.URL.Path == "/" {
			return CONTINUE, nil
		}

		if slices.Contains(proxy.Config.WhiteListPath, req.URL.Path) {
			return CONTINUE, nil
		}

		for _, whitelist := range proxy.Config.WhiteListDirs {
			if strings.HasPrefix(req.URL.Path, whitelist) {
				return CONTINUE, nil
			}
		}
	}

	return FINISHED, fmt.Errorf("failed to authenticate user")
}

func LogDirector(proxy *OchanocoProxy, req *http.Request, c *gin.Context) (bool, error) {
	request, err := httputil.DumpRequest(req, true)
	fmt.Printf("%v\n", string(request))
	err = makeError(err, "failed to dump headers to json: %v")
	logRawCommunication("", request, proxy)

	return CONTINUE, err
}
