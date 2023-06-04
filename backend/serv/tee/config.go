package tee

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"io/ioutil"
	"math/big"
	"time"

	"github.com/edgelesssys/ego/ecrypto"
)

const private_path = "./private.der"
const certificate_path = "./cert.der"

func ReadConfig() *tls.Config {
	bytes, err := ioutil.ReadFile(private_path)
	if err != nil {
		panic(err)
	}

	plaintext, err := ecrypto.Unseal(bytes, []byte{})

	if err != nil {
		panic(err)
	}

	priv, err := x509.ParsePKCS1PrivateKey(plaintext)

	if err != nil {
		panic(err)
	}

	cert, err := ioutil.ReadFile(certificate_path)
	if err != nil {
		panic(err)
	}

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

func InitConfig() *tls.Config {
	priv, _ := rsa.GenerateKey(rand.Reader, 2048)

	x509Priv := x509.MarshalPKCS1PrivateKey(priv)

	data, err := ecrypto.SealWithProductKey(x509Priv, []byte{})

	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(private_path, data, 0644)

	if err != nil {
		panic(err)
	}

	template := &x509.Certificate{
		SerialNumber: &big.Int{},
		Subject:      pkix.Name{CommonName: "localhost"},
		NotAfter:     time.Now().Add(time.Hour),
		DNSNames:     []string{"localhost"},
	}

	cert, _ := x509.CreateCertificate(rand.Reader, template, template, &priv.PublicKey, priv)

	err = ioutil.WriteFile(certificate_path, cert, 0644)

	tlsCfg := tls.Config{
		Certificates: []tls.Certificate{
			{
				Certificate: [][]byte{cert},
				PrivateKey:  priv,
			},
		},
	}

	token := setupAttestaion(&tlsCfg)
	fmt.Printf("token = %v\n", token)

	if err != nil {
		panic(err)
	}

	return &tlsCfg
}
