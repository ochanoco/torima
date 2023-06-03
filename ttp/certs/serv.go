package certs

import (
	"io"
	"log"
	"net/http"
)

const (
	DEFAULT_HOST string = "0.0.0.0:8080"
)

func SignPubkey(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		log.Printf("SignPubkey: Method %v is not allowed\n", req.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// verify logger installation before signing
	// ...

	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("SignPubkey: failed to read request body: %v\n", err)
		w.WriteHeader(500)
		return
	}

	signature, err := Sign(body)
	if err != nil {
		log.Printf("SignPubkey: failed to sign public key: %v\n", err)
		w.WriteHeader(500)
		return
	}

	signaturePem := EncodeSignaturePem(signature)
	w.Write(signaturePem)
}

func StartServer(addr string) {
	http.HandleFunc("/register", SignPubkey)
	http.ListenAndServe(addr, nil)
}
