package tee

import (
	"github.com/ochanoco/proxy/core"
	"fmt"
	"net/http"
)


const NAME = "tee"
const attestationProviderURL = "https://shareduks.uks.attest.azure.net"
const serverAddr = "0.0.0.0:8080"

func Run() *core.OchanocoProxy {
	secret := "testest"
	core.DEFAULT_PROXYWEB_PAGES = DEFAULT_PROXYWEB_PAGES

	core.SetupParsingUrl()

	proxyServ := core.ProxyServer(secret)
	return proxyServ
}

func Main() {	
	tlsConfig := setupTLS()

	proxyServ := Run()	
	setupAttestaion(proxyServ.Engine, tlsConfig)

	teeServer := http.Server{
		Addr:      serverAddr,
		TLSConfig: tlsConfig,
		Handler:   proxyServ.Engine,
	}

	err := teeServer.ListenAndServeTLS("", "")
	fmt.Println(err)
}
