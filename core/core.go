package core

import (
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

type TorimaPackageStatus = int

const (
	AuthNeeded TorimaPackageStatus = iota
	Authed
	NoAuthNeeded
	ForceStop
	Keep
)

type TorimaPackageTarget interface{ *http.Request | *http.Response }

type TorimaPackageContext[T TorimaPackageTarget] struct {
	Proxy         *TorimaProxy
	Target        T
	GinContext    *gin.Context
	PackageStatus TorimaPackageStatus
}

type TorimaDirectorPackageContext = TorimaPackageContext[*http.Request]
type TorimaModifyResponsePackageContext = TorimaPackageContext[*http.Response]

type TorimaDirector func(*TorimaDirectorPackageContext) (TorimaPackageStatus, error)
type TorimaModifyResponse func(*TorimaModifyResponsePackageContext) (TorimaPackageStatus, error)
type TorimaDirectors []func(*TorimaDirectorPackageContext) (TorimaPackageStatus, error)
type TorimaModifyResponses []func(*TorimaModifyResponsePackageContext) (TorimaPackageStatus, error)

type TorimaProxyWebPage = func(proxy *TorimaProxy, c *gin.RouterGroup)

type TorimaProxy struct {
	Directors       TorimaDirectors
	ModifyResponses TorimaModifyResponses
	ProxyWebPages   []TorimaProxyWebPage
	Engine          *gin.Engine
	Database        *Database
	ErrorHandler    *gin.HandlerFunc
	Config          *TorimaConfig
	RequestCount    int
}

func NewOchancoProxy(
	r *gin.Engine,
	directors TorimaDirectors,
	modifyResponses TorimaModifyResponses,
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
