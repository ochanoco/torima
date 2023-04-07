package core

import (
	"log"
	"net/url"
)

var ProxyRedirectUrl *url.URL
var ErrorUrl *url.URL
var AuthWebBaseUrl *url.URL
var ProxyWebBaseUrl *url.URL

func SetupParsingUrl() {
	var err error
	errorTemplate := "failed to parse url (%v)\n%v"

	if err != nil {
		log.Fatalf(errorTemplate, PROXY_REDIRECT_URL, err)
	}

	ErrorUrl, err = url.Parse(ERROR_URL)

	if err != nil {
		log.Fatalf(errorTemplate, ERROR_URL, err)
	}

	AuthWebBaseUrl, err = url.Parse(AUTHWEB_BASE)
	if err != nil {
		log.Fatalf(errorTemplate, AUTHWEB_BASE, err)
	}

	ProxyWebBaseUrl, err = url.Parse(PROXYWEB_BASE)
	if err != nil {
		log.Fatalf(errorTemplate, AUTHWEB_BASE, err)
	}

	ProxyRedirectUrl, err = url.Parse(PROXY_REDIRECT_URL)
	if err != nil {
		log.Fatalf(errorTemplate, AUTHWEB_BASE, err)
	}

}
