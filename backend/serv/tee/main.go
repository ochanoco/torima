package tee

import (
	"crypto/rsa"
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/ochanoco/proxy/core"
)

const NAME = "tee"
const attestationProviderURL = "https://shareduks.uks.attest.azure.net"
const serverAddr = "0.0.0.0:8080"

func Run(config *tls.Config) *core.OchanocoProxy {
	secret := "testest"

	rsaPriv := config.Certificates[0].PrivateKey.(*rsa.PrivateKey)

	core.DEFAULT_DIRECTORS = []core.OchanocoDirector{
		core.EnvRouteDirector,
		core.SetupLogVerifiableCommunicationDirector(rsaPriv),
	}

	core.DEFAULT_MODIFY_RESPONSES = []core.OchanocoModifyResponse{
		core.SetupLogVerifiableCommunicationResp(rsaPriv),
	}
	core.DEFAULT_PROXYWEB_PAGES = DEFAULT_PROXYWEB_PAGES

	core.SetupParsingUrl()

	proxyServ := core.ProxyServer(secret)
	return proxyServ
}

func Main() {
	tlsConfig := setupTLS()

	proxyServ := Run(tlsConfig)
	setupAttestaion(proxyServ.Engine, tlsConfig)

	teeServer := http.Server{
		Addr:      serverAddr,
		TLSConfig: tlsConfig,
		Handler:   proxyServ.Engine,
	}

	err := teeServer.ListenAndServeTLS("", "")
	fmt.Println(err)
}
