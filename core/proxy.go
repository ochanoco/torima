package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func runAllPackage[T TorimaPackageTarget](
	pkgs []func(*TorimaPackageContext[T]) (int, error),
	c *TorimaPackageContext[T]) {

	logger := NewFlowLogger()
	for _, pkg := range pkgs {
		status, err := pkg(c)
		logger.Add(pkg, status)

		c.PackageStatus = status

		if err != nil {
			abordGin(err, c.GinContext)
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

	LogReq(req)
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
		PackageStatus: Stay,
	}

	runAllPackage(proxy.ModifyResponses, &c)
	return nil
}
