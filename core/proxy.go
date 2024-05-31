package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ochanoco/torima/utils"
)

func runAllExtension[T TorimaPackageTarget](
	pkgs []func(*TorimaPackageContext[T]) (int, error),
	c *TorimaPackageContext[T]) {

	for _, pkg := range pkgs {
		status, err := pkg(c)

		if status != Keep {
			c.PackageStatus = status
		}

		if err != nil {
			utils.AbordGin(err, c.GinContext)
		}

		if status == ForceStop {
			break
		}
	}
}

/**
 * Directors is a list of functions that modify the
 * request before it is sent to the target server.
 **/
func (proxy *TorimaProxy) Director(req *http.Request, ginContext *gin.Context) {
	c := TorimaDirectorPackageContext{
		Proxy:         proxy,
		Target:        req,
		GinContext:    ginContext,
		PackageStatus: AuthNeeded,
	}

	runAllExtension[*http.Request](proxy.Directors, &c)
}

/**
  * ModifyResponses is a list of functions that modify the
  * response before it is sent to the client.
**/
func (proxy *TorimaProxy) ModifyResponse(res *http.Response, ginContext *gin.Context) error {
	c := TorimaModifyResponsePackageContext{
		Proxy:         proxy,
		Target:        res,
		GinContext:    ginContext,
		PackageStatus: Keep,
	}

	runAllExtension(proxy.ModifyResponses, &c)
	return nil
}
