package directors

import (
	"fmt"
	"strings"

	"github.com/ochanoco/torima/core"
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
