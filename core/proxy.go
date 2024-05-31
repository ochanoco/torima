package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ochanoco/torima/utils"
)

func runAllPackage[T TorimaPackageTarget](
	pkgs []func(*TorimaPackageContext[T]) (int, error),
	c *TorimaPackageContext[T]) {

	logger := utils.NewFlowLogger()
	for _, pkg := range pkgs {
		status, err := pkg(c)
		logger.Add(pkg, status)

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

	logger.Show()
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

	runAllPackage[*http.Request](proxy.Directors, &c)

	utils.LogReq(req)
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

	runAllPackage(proxy.ModifyResponses, &c)
	return nil
}
