package directors

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/ochanoco/ninsho"
	gin_ninsho "github.com/ochanoco/ninsho/extension/gin"
	"github.com/ochanoco/torima/core"
	"github.com/ochanoco/torima/utils"

	"golang.org/x/exp/slices"
)

func RouteDirector(host string, c *core.TorimaDirectorPackageContext) (core.TorimaPackageStatus, error) {
	c.Target.URL.Host = host

	// just to be sure
	c.Target.Header.Del("X-Torima-Proxy-Token")
	c.Target.Header.Set("X-Torima-Proxy-Token", core.SECRET)

	c.Target.URL.Scheme = c.Proxy.Config.Scheme

	return core.Keep, nil
}

func DefaultRouteDirector(c *core.TorimaDirectorPackageContext) (core.TorimaPackageStatus, error) {
	if strings.HasPrefix(c.Target.URL.Path, "/torima/") {
		return core.Keep, nil
	}

	host := c.Proxy.Config.DefaultOrigin

	if host == "" {
		err := fmt.Errorf("failed to get destination config (%s)", host)
		return core.ForceStop, err
	}

	return RouteDirector(host, c)
}

func ThirdPartyDirector(c *core.TorimaDirectorPackageContext) (core.TorimaPackageStatus, error) {
	path := strings.Split(c.Target.URL.Path, "/")
	hasRedirectPrefix := strings.HasPrefix(c.Target.URL.Path, "/torima/redirect/")

	if !hasRedirectPrefix || len(path) < 3 {
		return core.Keep, nil
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

	return core.Keep, nil
}

func SanitizeHeaderDirector(c *core.TorimaDirectorPackageContext) (core.TorimaPackageStatus, error) {
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

	return core.Keep, nil

}

func SkipAuthDirector(c *core.TorimaDirectorPackageContext) (core.TorimaPackageStatus, error) {
	if c.Target.Method == "GET" && c.Target.URL.RawQuery == "" {
		if c.Target.URL.Path == "/" {
			return core.NoAuthNeeded, nil
		}

		if slices.Contains(c.Proxy.Config.SkipAuthList, c.Target.URL.Path) {
			return core.NoAuthNeeded, nil
		}
	}

	return core.AuthNeeded, nil
}

func ForceAuthDirector(c *core.TorimaDirectorPackageContext) (core.TorimaPackageStatus, error) {
	if slices.Contains(c.Proxy.Config.ForceAuthList, c.Target.URL.Path) {
		return core.AuthNeeded, nil
	}

	return core.Keep, nil
}

func AuthDirector(c *core.TorimaDirectorPackageContext) (core.TorimaPackageStatus, error) {
	if c.PackageStatus == core.NoAuthNeeded {
		return core.NoAuthNeeded, nil
	}

	user, err := gin_ninsho.LoadUser[ninsho.LINE_USER](c.GinContext)

	// just to be sure
	c.Target.Header.Del("X-Torima-UserID")

	if err != nil {
		err = utils.MakeError(err, "failed to get user from session: ")
		return core.ForceStop, err
	}

	if user != nil {
		c.Target.Header.Set("X-Torima-UserID", user.Sub)
		return core.Authed, nil
	}

	return core.ForceStop, utils.MakeError(fmt.Errorf(""), utils.UnauthorizedErrorTag)
}

func MakeLogDirector(flag string) core.TorimaDirector {
	return func(c *core.TorimaDirectorPackageContext) (core.TorimaPackageStatus, error) {
		request, err := httputil.DumpRequest(c.Target, true)

		if err != nil {
			err = utils.MakeError(err, "failed to dump headers to json: ")
			return core.ForceStop, err
		}

		splited := bytes.Split(request, []byte("\r\n\r\n"))

		header := splited[0]
		headerLen := len(header)

		body := request[headerLen:]

		l := c.Proxy.Database.CreateRequestLog(string(header), body, flag)
		_, err = l.Save(c.Proxy.Database.Ctx)

		if err != nil {
			err = utils.MakeError(err, "failed to save request: ")
			return core.ForceStop, err
		}

		return core.Keep, err
	}
}

var BeforeLogDirector = MakeLogDirector("before")
var AfterLogDirector = MakeLogDirector("after")
