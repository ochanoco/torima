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

func RouteDirector(host string, proxy *TorimaProxy, req *http.Request, c *gin.Context) (bool, error) {
	req.URL.Host = host

	// just to be sure
	req.Header.Del("X-Torima-Proxy-Token")
	req.Header.Set("X-Torima-Proxy-Token", SECRET)

	req.URL.Scheme = proxy.Config.Scheme

	return CONTINUE, nil
}

func DefaultRouteDirector(proxy *TorimaProxy, req *http.Request, c *gin.Context) (bool, error) {
	if strings.HasPrefix(req.URL.Path, "/torima/") {
		return CONTINUE, nil
	}

	host := proxy.Config.DefaultOrigin

	if host == "" {
		err := fmt.Errorf("failed to get destination config (%s)", host)
		return FINISHED, err
	}

	return RouteDirector(host, proxy, req, c)
}

func ThirdPartyDirector(proxy *TorimaProxy, req *http.Request, c *gin.Context) (bool, error) {
	path := strings.Split(req.URL.Path, "/")
	hasRedirectPrefix := strings.HasPrefix(req.URL.Path, "/torima/redirect/")

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

func SanitizeHeaderDirector(proxy *TorimaProxy, req *http.Request, c *gin.Context) (bool, error) {
	headers := http.Header{
		"Host":       {proxy.Config.Host},
		"User-Agent": {"torima"},

		"Content-Type":   req.Header["Content-Type"],
		"Content-Length": req.Header["Content-Length"],

		"Accept":     req.Header["Accept"],
		"Connection": req.Header["Connection"],

		"Accept-Encoding": req.Header["Accept-Encoding"],
		"Accept-Language": req.Header["Accept-Language"],

		"Cookie": req.Header["Cookie"],
	}

	req.Header = headers

	return CONTINUE, nil

}

func AuthDirector(proxy *TorimaProxy, req *http.Request, c *gin.Context) (bool, error) {
	user, err := gin_ninsho.LoadUser[ninsho.LINE_USER](c)

	// just to be sure
	req.Header.Del("X-Torima-UserID")

	if err != nil {
		err = makeError(err, "failed to get user from session: ")
		return FINISHED, err
	}

	if user != nil {
		req.Header.Set("X-Torima-UserID", user.Sub)
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

func MakeLogDirector(flag string) TorimaDirector {
	return func(proxy *TorimaProxy, req *http.Request, c *gin.Context) (bool, error) {
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
