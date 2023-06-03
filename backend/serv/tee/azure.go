package tee

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"log"
	"math/big"
	"time"

	"github.com/edgelesssys/ego/enclave"
	"github.com/gin-gonic/gin"
)

func setupTLS() *tls.Config {
	template := &x509.Certificate{
		SerialNumber: &big.Int{},
		Subject:      pkix.Name{CommonName: "localhost"},
		NotAfter:     time.Now().Add(time.Hour),
		DNSNames:     []string{"localhost"},
	}

	priv, _ := rsa.GenerateKey(rand.Reader, 2048)
	cert, _ := x509.CreateCertificate(rand.Reader, template, template, &priv.PublicKey, priv)

	tlsCfg := tls.Config{
		Certificates: []tls.Certificate{
			{
				Certificate: [][]byte{cert},
				PrivateKey:  priv,
			},
		},
	}
	return &tlsCfg
}

func setupAttestaion(r *gin.Engine, tlsConfig *tls.Config) {
	token, err := enclave.CreateAzureAttestationToken(tlsConfig.Certificates[0].Certificate[0], attestationProviderURL)
	if err != nil {
		log.Print("Run without attestation!!!!!\n")
	} else {
		log.Print("Created an Microsoft Azure Attestation Token.")
	}

	r.Any("/token", func(c *gin.Context) {
		c.String(200, token)
	})
}