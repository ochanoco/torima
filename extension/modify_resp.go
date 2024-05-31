package extension

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/ochanoco/torima/core"
	"github.com/ochanoco/torima/utils"
)

func MainModifyResponse(proxy *core.TorimaProxy, res *http.Response) {
	fmt.Printf("=> %v\n", res.Request.URL)
}

func InjectHTML(html string, c *core.TorimaModifyResponsePackageContext) (core.TorimaPackageStatus, error) {
	document, err := goquery.NewDocumentFromReader(c.Target.Body)
	if err != nil {
		log.Fatal(err)
	}

	document.Find("body").AppendHtml(html)

	html, err = document.Html()
	if err != nil {
		return core.ForceStop, err
	}

	// fmt.Printf("%v", html)

	b := []byte(html)
	c.Target.Body = io.NopCloser(bytes.NewReader(b))

	c.Target.Header.Set("Content-Length", strconv.Itoa(len(b)))
	c.Target.ContentLength = int64(len(b))

	return core.Keep, nil
}

func InjectServiceWorkerModifyResponse(c *core.TorimaModifyResponsePackageContext) (core.TorimaPackageStatus, error) {
	contentType := c.Target.Header.Get("Content-Type")

	if contentType != "text/html; charset=utf-8" {
		return core.Keep, nil
	}

	html := utils.Scripts + "\n"

	return InjectHTML(html, c)
}
