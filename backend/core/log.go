package core

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ochanoco/proxy/ent"
)

func logToDB(header http.Header, body io.ReadCloser, proxy *OchanocoProxy, c *gin.Context) (*ent.ServiceLog, error) {
	time := time.Now()

	headerJson, err := DumpHeader(header)
	if err != nil {
		return nil, fmt.Errorf("failed to dump headers to json: %v", err)
	}

	bodyBuf, err := io.ReadAll(body)
	if err != nil {
		return nil, fmt.Errorf("non-nil error while reading request body: %v", err)
	}

	body.Close()
	io.NopCloser(bytes.NewBuffer(bodyBuf))

	l := proxy.Database.CreateServiceLog(time, headerJson, bodyBuf)
	result, err := proxy.Database.SaveServiceLog(l)

	if err != nil {
		return nil, fmt.Errorf("ResponseLogDirector: failed to save: %v", err)
	}

	return result, nil
}
