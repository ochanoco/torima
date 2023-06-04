package core

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ochanoco/proxy/ent"
)

func logRaw(header string, body []byte, proxy *OchanocoProxy) (*ent.ServiceLog, error) {
	time := time.Now()

	l := proxy.Database.CreateServiceLog(time, header, body)
	return proxy.Database.SaveServiceLog(l)
}

func logCommunication(header http.Header, body *io.ReadCloser, proxy *OchanocoProxy) (*ent.ServiceLog, error) {
	headerJson, err := DumpHeader(header)
	if err != nil {
		return nil, fmt.Errorf("failed to dump headers to json: %v", err)
	}

	// There are kinds of methods which does not have bodies (i.e., GET, HEAD, OPTIONS, TRACE).
	if *body == nil {
		return logRaw(headerJson, nil, proxy)
	}

	bodyBuf, err := ReadHTTPBody(body)
	if body == nil {
		return nil, err
	}

	return logRaw(headerJson, bodyBuf, proxy)
}

func logRawHashChain(new, previous []byte, proxy *OchanocoProxy, privkey *rsa.PrivateKey) (*ent.HashChain, error) {
	hasher := crypto.SHA256.New()
	hasher.Write(previous)
	hasher.Write(new)
	hash := hasher.Sum(nil)

	signature, err := rsa.SignPSS(rand.Reader, privkey, crypto.SHA3_256, hash, &rsa.PSSOptions{
		SaltLength: rsa.PSSSaltLengthAuto,
	})

	if err != nil {
		return nil, err
	}

	chainCreate := proxy.Database.CreateHashChain(hash, signature)
	return proxy.Database.SaveHashChain(chainCreate)

}

func logVerifiableCommunication(header http.Header, body *io.ReadCloser, proxy *OchanocoProxy, privkey *rsa.PrivateKey) (*ent.ServiceLog, *ent.HashChain, error) {
	last, err := proxy.Database.FindLastHashChain()

	if err != nil {
		last, err = logRawHashChain([]byte("this may be first hash"), []byte{}, proxy, privkey)

		if err != nil {
			return nil, nil, err
		}
	}

	log, err := logCommunication(header, body, proxy)

	if err != nil {
		return nil, nil, err
	}

	new := append([]byte(log.Headers), log.Body...)

	// todo: change timing to log
	chain, err := logRawHashChain(new, last.Hash, proxy, privkey)

	if err != nil {
		return nil, nil, err
	}

	return log, chain, nil
}

func SetupLogVerifiableCommunicationDirector(privkey *rsa.PrivateKey) OchanocoDirector {
	return func(proxy *OchanocoProxy, req *http.Request, c *gin.Context) bool {
		fmt.Printf("body1: %v\n", req.Body)

		_, _, err := logVerifiableCommunication(req.Header, &req.Body, proxy, privkey)

		if err != nil {
			fmt.Printf("LogModifyResponse: %v\n", err)
			return FINISHED
		}

		return CONTINUE
	}
}

func SetupLogVerifiableCommunicationResp(privkey *rsa.PrivateKey) OchanocoModifyResponse {
	return func(proxy *OchanocoProxy, resp *http.Response, c *gin.Context) bool {
		_, _, err := logVerifiableCommunication(resp.Header, &resp.Body, proxy, privkey)

		if err != nil {
			fmt.Printf("LogModifyResponse: %v\n", err)
			return FINISHED
		}

		return CONTINUE
	}
}
