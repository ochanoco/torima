package cloud

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ochanoco/proxy/core"
)

func CloudRouteDirector(proxy *core.OchanocoProxy, req *http.Request, c *gin.Context) bool {
	project, err := FindServiceProviderByHost(proxy.Database, req.Host)

	if err != nil {
		msg := fmt.Sprintf("failed to get destination site (%s)", req.Host)
		core.GoToErrorPage(msg, err, req)
		return core.FINISHED
	}

	return core.RouteDirector(project.DestinationIP, proxy, req, c)
}
