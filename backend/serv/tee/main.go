package tee

import (
	"crypto/rsa"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"

	"github.com/ochanoco/proxy/core"
)

const NAME = "tee"
const attestationProviderURL = "https://shareduks.uks.attest.azure.net"
const serverAddr = "0.0.0.0:8080"

func Run(rsaPriv *rsa.PrivateKey) *core.OchanocoProxy {
	secret := "testest"

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
	_, err1 := os.Stat(private_path)
	_, err2 := os.Stat(certificate_path)

	hasKeyExist := err1 == nil && err2 == nil

	var tlsCfg *tls.Config

	if hasKeyExist {
		tlsCfg = ReadConfig()
	} else {
		tlsCfg = InitConfig()
	}

	priv := tlsCfg.Certificates[0].PrivateKey.(*rsa.PrivateKey)
	proxyServ := Run(priv)

	err := verifyDB(&priv.PublicKey, proxyServ)
	if err != nil {
		fmt.Printf("verifyDB: %v\n", err)
		return
	}

	teeServer := http.Server{
		Addr:      serverAddr,
		TLSConfig: tlsCfg,
		Handler:   proxyServ.Engine,
	}

	fmt.Printf("tee server is listening on %v\n", serverAddr)

	err = teeServer.ListenAndServeTLS("", "")
	fmt.Println(err)
}
