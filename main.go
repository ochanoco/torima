package main

import (
	"net/http/httptest"
	"net/http/httputil"
)

func main() {
	rp := httputil.ReverseProxy{
		Director:       director,
		ModifyResponse: modifyResponse,
	}

	serv := httptest.NewServer(&rp)
	defer serv.Close()

	serv.Config.ListenAndServe()
}
