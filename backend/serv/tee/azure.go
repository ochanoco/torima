package tee

import (
	"crypto/tls"
	"log"

	"github.com/edgelesssys/ego/enclave"
)

func setupAttestaion(tlsConfig *tls.Config) string {
	token, err := enclave.CreateAzureAttestationToken(tlsConfig.Certificates[0].Certificate[0], attestationProviderURL)
	if err != nil {
		log.Print("Run without attestation!!!!!\n")
	} else {
		log.Print("Created an Microsoft Azure Attestation Token.")
	}

	return token
}
