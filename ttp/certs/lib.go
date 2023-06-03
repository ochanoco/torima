package certs

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

func EncodePem(data []byte, typ string) []byte {
	block := &pem.Block{
		Type:  typ,
		Bytes: data,
	}
	buff := &bytes.Buffer{}
	pem.Encode(buff, block)
	return buff.Bytes()
}

func EncodePubkeyPem(pubkey *rsa.PublicKey) []byte {
	typ := "RSA PUBLIC KEY"
	data := x509.MarshalPKCS1PublicKey(pubkey)
	return EncodePem(data, typ)
}

func EncodeSignaturePem(signature []byte) []byte {
	typ := "CERTIFICATE"
	return EncodePem(signature, typ)
}
