package core

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ochanoco/proxy/ent"
)

func logRaw(header string, body []byte, proxy *OchanocoProxy) (*ent.ServiceLog, error) {
	time := time.Now()

	l := proxy.Database.CreateServiceLog(time, header, body)
	return proxy.Database.SaveServiceLog(l)
}

func LogCommunication(header http.Header, body *io.ReadCloser, proxy *OchanocoProxy) (*ent.ServiceLog, error) {
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
