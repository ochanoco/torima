package directors

import (
	"github.com/ochanoco/ninsho"
	gin_ninsho "github.com/ochanoco/ninsho/extension/gin"
	"github.com/ochanoco/torima/core"
	"github.com/ochanoco/torima/utils"

	"golang.org/x/exp/slices"
)

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

func AuthDirector(c *core.TorimaDirectorPackageContext) (core.TorimaPackageStatus, error) {
	user, err := gin_ninsho.LoadUser[ninsho.LINE_USER](c.GinContext)
	// just to be sure
	c.Target.Header.Del("X-Torima-UserID")

	if err != nil || user == nil {
		if c.PackageStatus == core.NoAuthNeeded {
			return core.NoAuthNeeded, nil
		}

		err = utils.MakeError(err, "failed to authenticate: ")
		return core.ForceStop, err
	}

	c.Target.Header.Set("X-Torima-UserID", user.Sub)
	return core.Authed, nil
}
