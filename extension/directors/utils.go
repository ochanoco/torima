package directors

import (
	"net/http"
	"strings"

	"github.com/ochanoco/torima/core"
)

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
			return BasicRoute(origin, c)
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
