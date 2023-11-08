package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const CONTINUE = true
const FINISHED = false

type TorimaPackageArgument interface{ *http.Request | *http.Response }

func runAllPackage[T TorimaPackageArgument](
	pkgs []func(*TorimaProxy, T, *gin.Context) (bool, error),
	args T, proxy *TorimaProxy, c *gin.Context) {

	logger := NewFlowLogger()

	for _, pkg := range pkgs {
		isContinuing, err := pkg(proxy, args, c)
		logger.Add(pkg, isContinuing)

		if err != nil {
			abordGin(proxy, err, c)
		}

		if !isContinuing {
			break
		}
	}

	logger.Show()
}

/**
 * Directors is a list of functions that modify the
 * request before it is sent to the target server.
 **/
func (proxy *TorimaProxy) Director(req *http.Request, c *gin.Context) {
	runAllPackage(proxy.Directors, req, proxy, c)

	LogReq(req)
}

/**
  * ModifyResponses is a list of functions that modify the
  * response before it is sent to the client.
**/
func (proxy *TorimaProxy) ModifyResponse(res *http.Response, c *gin.Context) error {
	runAllPackage(proxy.ModifyResponses, res, proxy, c)
	return nil
}
