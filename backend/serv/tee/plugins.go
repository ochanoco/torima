package tee

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ochanoco/proxy/core"
	"github.com/ochanoco/proxy/ent"
)

type HashChain = ent.HashChain

func CalcHashChain(log *ent.ServiceLog, previous *ent.HashChain) []byte {
	var previous_buf []byte

	if previous_buf == nil {
		previous_buf = []byte("invalid")
	} else {
		previous_buf = previous.Hash
	}

	hasher := crypto.SHA256.New()
	hasher.Write([]byte(log.Headers))
	hasher.Write(log.Body)
	hasher.Write(previous_buf)
	hash := hasher.Sum(nil)

	return hash
}

func logRawHashChain(new *ent.ServiceLog, previous *ent.HashChain, proxy *core.OchanocoProxy, privkey *rsa.PrivateKey) (*ent.HashChain, error) {
	hash := CalcHashChain(new, previous)

	signature, err := rsa.SignPSS(rand.Reader, privkey, crypto.SHA3_256, hash, &rsa.PSSOptions{
		SaltLength: rsa.PSSSaltLengthAuto,
	})

	if err != nil {
		return nil, err
	}

	chainCreate := proxy.Database.CreateHashChain(hash, signature)
	return proxy.Database.SaveHashChain(chainCreate)

}

func logVerifiableCommunication(header http.Header, body *io.ReadCloser, proxy *core.OchanocoProxy, privkey *rsa.PrivateKey) (*ent.ServiceLog, *ent.HashChain, error) {
	last, err := proxy.Database.FindLastHashChain()

	if err != nil {
		return nil, nil, err
	}

	log, err := core.LogCommunication(header, body, proxy)

	if err != nil {
		return nil, nil, err
	}

	// todo: change timing to log
	chain, err := logRawHashChain(log, last, proxy, privkey)

	if err != nil {
		return nil, nil, err
	}

	return log, chain, nil
}

func SetupLogVerifiableCommunicationDirector(privkey *rsa.PrivateKey) core.OchanocoDirector {
	return func(proxy *core.OchanocoProxy, req *http.Request, c *gin.Context) bool {
		_, _, err := logVerifiableCommunication(req.Header, &req.Body, proxy, privkey)

		if err != nil {
			fmt.Printf("LogModifyResponse: %v\n", err)
			return core.FINISHED
		}

		return core.CONTINUE
	}
}

func SetupLogVerifiableCommunicationResp(privkey *rsa.PrivateKey) core.OchanocoModifyResponse {
	return func(proxy *core.OchanocoProxy, resp *http.Response, c *gin.Context) bool {
		_, _, err := logVerifiableCommunication(resp.Header, &resp.Body, proxy, privkey)

		if err != nil {
			fmt.Printf("LogModifyResponse: %v\n", err)
			return core.FINISHED
		}

		return core.CONTINUE
	}
}
