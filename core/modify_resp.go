package core

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

func MainModifyResponse(proxy *TorimaProxy, res *http.Response) {
	fmt.Printf("=> %v\n", res.Request.URL)
}

func InjectHTMLModifyResponse(html string, proxy *TorimaProxy, res *http.Response, c *gin.Context) (bool, error) {
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	document.Find("body").AppendHtml(html)

	html, err = document.Html()
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

func InjectServiceWorkerModifyResponse(proxy *TorimaProxy, res *http.Response, c *gin.Context) (bool, error) {
	contentType := res.Header.Get("Content-Type")

	if contentType != "text/html; charset=utf-8" {
		return CONTINUE, nil
	}

	html := scripts + "\n"

	return InjectHTMLModifyResponse(html, proxy, res, c)
}
