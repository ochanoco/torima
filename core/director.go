package core

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/ochanoco/ninsho"
	gin_ninsho "github.com/ochanoco/ninsho/extension/gin"

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
	user, err := gin_ninsho.GetUser[ninsho.LINE_USER](c)

	if err != nil {
		err = makeError(err, "failed to get user from session: ")
		return FINISHED, err
	}

	if user != nil {
		req.Header.Set("X-Ochanoco-UserID", user.Sub)
		return CONTINUE, nil
	}

	if req.Method == "GET" && req.URL.RawQuery == "" {
		if req.URL.Path == "/" {
			return CONTINUE, nil
		}

		if slices.Contains(proxy.Config.WhiteListPath, req.URL.Path) {
			return CONTINUE, nil
		}
	}

	return FINISHED, makeError(fmt.Errorf(""), unauthorizedErrorTag)
}

func MakeLogDirector(flag string) OchanocoDirector {
	return func(proxy *OchanocoProxy, req *http.Request, c *gin.Context) (bool, error) {
		request, err := httputil.DumpRequest(req, true)

		if err != nil {
			err = makeError(err, "failed to dump headers to json: ")
			return FINISHED, err
		}

		splited := bytes.Split(request, []byte("\r\n\r\n"))

		header := splited[0]
		headerLen := len(header)

		body := request[headerLen:]

		l := proxy.Database.CreateRequestLog(string(header), body, flag)
		_, err = l.Save(proxy.Database.Ctx)

		if err != nil {
			err = makeError(err, "failed to save request: ")
			return FINISHED, err
		}

		return CONTINUE, err
	}
}

var BeforeLogDirector = MakeLogDirector("before")
var AfterLogDirector = MakeLogDirector("after")
