package main

import (
	"net/http"
	"net/http/httputil"
)

func main() {
	rp := httputil.ReverseProxy{
		Director:       director,
		ModifyResponse: modifyResponse,
	}

	serv := http.Server{
		Addr:    ":9000",
		Handler: &rp,
	}

	serv.ListenAndServe()
}
