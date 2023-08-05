package core

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

func MainModifyResponse(proxy *OchanocoProxy, res *http.Response) {
	fmt.Printf("=> %v\n", res.Request.URL)
}

func LogModifyResponse(proxy *OchanocoProxy, res *http.Response, c *gin.Context) (bool, error) {
	response, err := httputil.DumpResponse(res, true)
	fmt.Printf("%v\n", string(response))

	err = makeError(err, "failed to dump response: %v")
	logRawCommunication("response", "", response, proxy)

	return CONTINUE, err
}
