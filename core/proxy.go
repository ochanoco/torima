package core

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const CONTINUE = true
const FINISHED = false

type OchanocoPackageArgument interface{ *http.Request | *http.Response }

func runAllPackage[T OchanocoPackageArgument](
	pkgs []func(*OchanocoProxy, T, *gin.Context) (bool, error),
	args T, proxy *OchanocoProxy, c *gin.Context) {

	logger := NewFlowLogger()

	for _, pkg := range pkgs {
		isContinuing, err := pkg(proxy, args, c)

		logger.Add(pkg, isContinuing)

		if err != nil {
			panic(err)
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
func (proxy *OchanocoProxy) Director(req *http.Request, c *gin.Context) {
	req.URL.Scheme = "http"

	runAllPackage(proxy.Directors, req, proxy, c)
	LogReq(req)
}

/**
  * ModifyResponses is a list of functions that modify the
  * response before it is sent to the client.
**/
func (proxy *OchanocoProxy) ModifyResponse(res *http.Response, c *gin.Context) error {
	runAllPackage(proxy.ModifyResponses, res, proxy, c)
	return nil
}

func errorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.ByType(gin.ErrorTypePublic).Last()
		if err == nil {
			return
		}

		log.Print(err.Err)

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	}
}
