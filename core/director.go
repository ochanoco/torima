package core

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/ochanoco/ninsho"
	gin_ninsho "github.com/ochanoco/ninsho/extension/gin"

	"golang.org/x/exp/slices"
)

func RouteDirector(host string, c *TorimaDirectorPackageContext) (TorimaPackageStatus, error) {
	c.Target.URL.Host = host

	// just to be sure
	c.Target.Header.Del("X-Torima-Proxy-Token")
	c.Target.Header.Set("X-Torima-Proxy-Token", SECRET)

	c.Target.URL.Scheme = c.Proxy.Config.Scheme

	return Stay, nil
}

func DefaultRouteDirector(c *TorimaDirectorPackageContext) (TorimaPackageStatus, error) {
	if strings.HasPrefix(c.Target.URL.Path, "/torima/") {
		return Stay, nil
	}

	host := c.Proxy.Config.DefaultOrigin

	if host == "" {
		err := fmt.Errorf("failed to get destination config (%s)", host)
		return ForceStop, err
	}

	return RouteDirector(host, c)
}

func ThirdPartyDirector(c *TorimaDirectorPackageContext) (TorimaPackageStatus, error) {
	path := strings.Split(c.Target.URL.Path, "/")
	hasRedirectPrefix := strings.HasPrefix(c.Target.URL.Path, "/torima/redirect/")

	if !hasRedirectPrefix || len(path) < 3 {
		return Stay, nil
	}

	for _, origin := range c.Proxy.Config.ProtectionScope {
		if origin == path[3] {
			c.Target.Host = origin
			c.Target.URL.Host = origin

			p := strings.Join(path[4:], "/")
			c.Target.URL.Path = "/" + p

			c.Target.URL.Scheme = "https"
			return RouteDirector(origin, c)
		}
	}

	return Stay, nil
}

func SanitizeHeaderDirector(c *TorimaDirectorPackageContext) (TorimaPackageStatus, error) {
	headers := http.Header{
		"Host":       {c.Proxy.Config.Host},
		"User-Agent": {"torima"},

		"Content-Type":   c.Target.Header["Content-Type"],
		"Content-Length": c.Target.Header["Content-Length"],

		"Accept":     c.Target.Header["Accept"],
		"Connection": c.Target.Header["Connection"],

		"Accept-Encoding": c.Target.Header["Accept-Encoding"],
		"Accept-Language": c.Target.Header["Accept-Language"],

		"Cookie": c.Target.Header["Cookie"],
	}

	c.Target.Header = headers

	return Stay, nil

}

func AuthDirector(c *TorimaDirectorPackageContext) (TorimaPackageStatus, error) {
	user, err := gin_ninsho.LoadUser[ninsho.LINE_USER](c.GinContext)

	// just to be sure
	c.Target.Header.Del("X-Torima-UserID")

	if err != nil {
		err = makeError(err, "failed to get user from session: ")
		return ForceStop, err
	}

	if user != nil {
		c.Target.Header.Set("X-Torima-UserID", user.Sub)
		return Authed, nil
	}

	if c.Target.Method == "GET" && c.Target.URL.RawQuery == "" {
		if c.Target.URL.Path == "/" {
			return NoAuthNeeded, nil
		}

		if slices.Contains(c.Proxy.Config.WhiteListPath, c.Target.URL.Path) {
			return NoAuthNeeded, nil
		}
	}

	return AuthNeeded, makeError(fmt.Errorf(""), unauthorizedErrorTag)
}

func MakeLogDirector(flag string) TorimaDirector {
	return func(c *TorimaDirectorPackageContext) (TorimaPackageStatus, error) {
		request, err := httputil.DumpRequest(c.Target, true)

		if err != nil {
			err = makeError(err, "failed to dump headers to json: ")
			return ForceStop, err
		}

		splited := bytes.Split(request, []byte("\r\n\r\n"))

		header := splited[0]
		headerLen := len(header)

		body := request[headerLen:]

		l := c.Proxy.Database.CreateRequestLog(string(header), body, flag)
		_, err = l.Save(c.Proxy.Database.Ctx)

		if err != nil {
			err = makeError(err, "failed to save request: ")
			return ForceStop, err
		}

		return Stay, err
	}
}

var BeforeLogDirector = MakeLogDirector("before")
var AfterLogDirector = MakeLogDirector("after")
