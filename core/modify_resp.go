package core

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

func MainModifyResponse(proxy *OchanocoProxy, res *http.Response) {
	fmt.Printf("=> %v\n", res.Request.URL)
}

func LogModifyResponse(proxy *OchanocoProxy, res *http.Response, c *gin.Context) (bool, error) {
	response, err := httputil.DumpResponse(res, true)
	fmt.Printf("%v\n", string(response))

	err = makeError(err, "failed to dump response: ")
	logRawCommunication("response", response, proxy)

	return CONTINUE, err
}

func InjectResponse(proxy *OchanocoProxy, res *http.Response, c *gin.Context) (bool, error) {
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	document.Find("body").AppendHtml("<script>alert('hello')</script>")

	html, err := document.Html()
	if err != nil {
		return FINISHED, err
	}

	fmt.Printf("%v", html)

	b := []byte(html)
	res.Body = ioutil.NopCloser(bytes.NewReader(b))

	res.Header.Set("Content-Length", strconv.Itoa(len(b)))
	res.ContentLength = int64(len(b))

	return CONTINUE, nil
}
