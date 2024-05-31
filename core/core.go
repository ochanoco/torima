package core

import (
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

type TorimaDirector = func(proxy *TorimaProxy, req *http.Request, c *gin.Context) (bool, error)
type TorimaModifyResponse = func(proxy *TorimaProxy, resp *http.Response, c *gin.Context) (bool, error)
type TorimaProxyWebPage = func(proxy *TorimaProxy, c *gin.RouterGroup)

type TorimaProxy struct {
	Directors       []TorimaDirector
	ModifyResponses []TorimaModifyResponse
	ProxyWebPages   []TorimaProxyWebPage
	Engine          *gin.Engine
	Database        *Database
	ErrorHandler    *gin.HandlerFunc
	Config          *TorimaConfig
	RequestCount    int
}

func NewOchancoProxy(
	r *gin.Engine,
	directors []TorimaDirector,
	modifyResponses []TorimaModifyResponse,
	proxyWebPages []TorimaProxyWebPage,
	config *TorimaConfig,
	database *Database,
) TorimaProxy {
	proxy := TorimaProxy{}

	proxy.Directors = directors
	proxy.ModifyResponses = modifyResponses

	proxy.ProxyWebPages = proxyWebPages
	proxy.Database = database

	proxy.Engine = r
	proxy.Config = config

	specialPath := r.Group("/torima")
	for _, webPage := range proxy.ProxyWebPages {
		webPage(&proxy, specialPath)
	}

	r.NoRoute(func(c *gin.Context) {
		director := func(req *http.Request) {
			proxy.Director(req, c)
		}

		modifyResp := func(resp *http.Response) error {
			return proxy.ModifyResponse(resp, c)
		}

		proxy := httputil.ReverseProxy{
			Director:       director,
			ModifyResponse: modifyResp,
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	})

	return proxy
}
