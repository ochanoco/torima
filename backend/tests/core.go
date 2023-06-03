package tests

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ochanoco/proxy/core"
	"github.com/ochanoco/proxy/serv/line"
)

const TEST_RESP_BODY1 = "hoge"
const TEST_RESP_BODY2 = "piyo"
const TEST_RESP_ERROR_BODY = "error"
const TEST_RESP_REDIRECT_BODY = "redirect"

func ParseURL(t *testing.T, URL string) *url.URL {
	url, err := url.Parse(URL)

	if err != nil {
		t.Fatalf("parse: %v", err)
	}

	return url
}

type ProxyTester interface {
	Start(t *testing.T, proxy *core.OchanocoProxy, proxyServ *httptest.Server, testServ *httptest.Server)

	Directors(t *testing.T, URL string) []core.OchanocoDirector
	ModifyResps(t *testing.T) []core.OchanocoModifyResponse
	TestServers(t *testing.T) (*httptest.Server, *httptest.Server, *httptest.Server)

	Request(t *testing.T, url string) *http.Response
	Check(t *testing.T, resp *http.Response)
}

func RunCommonTest(t *testing.T, tester ProxyTester, name string) {
	var proxy *core.OchanocoProxy
	testServ, errorServ, redirectServ := tester.TestServers(t)
	core.ERROR_URL = errorServ.URL
	core.PROXY_REDIRECT_URL = redirectServ.URL

	directors := tester.Directors(t, testServ.URL)
	modifyResps := tester.ModifyResps(t)

	core.DEFAULT_DIRECTORS = directors
	core.DEFAULT_MODIFY_RESPONSES = modifyResps

	proxy = line.Run()
	proxyServ := httptest.NewServer(proxy.Engine)

	tester.Start(t, proxy, proxyServ, testServ)

	t.Run(name, func(t *testing.T) {
		resp := tester.Request(t, proxyServ.URL)

		tester.Check(t, resp)
	})
}

/////////

func RequestGetforTest(t *testing.T, url string) *http.Response {
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		t.Fatal(err)
	}

	resp, err := new(http.Client).Do(req)
	if err != nil {
		t.Fatal(err)
	}

	return resp
}

func RequestPostWithCookieForTest(t *testing.T, url, cookie string) *http.Response {
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		t.Error(err)
	}

	req.Header.Set("Cookie", cookie)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		t.Error(err)
	}

	return resp
}

func CheckResponseWithBody(t *testing.T, resp *http.Response, expect string) {
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("%v", err)
	}

	respStr := string(respBytes)

	if respStr != expect {
		t.Fatalf("wrong response: '%s' != '%s'(expected)", respStr, expect)
	}
}

func MakeSimpleServers() (*httptest.Server, *httptest.Server, *httptest.Server) {
	s := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		fmt.Fprint(writer, TEST_RESP_BODY1)
	}))

	return s, s, s
}

func MakesSimpleDirectors(t *testing.T, URL string) []core.OchanocoDirector {
	test := func(proxy *core.OchanocoProxy, req *http.Request, c *gin.Context) bool {
		url := ParseURL(t, URL)

		req.URL.Scheme = url.Scheme
		req.URL.Host = url.Host
		req.URL.Path = "/"

		return core.FINISHED
	}

	return []core.OchanocoDirector{test}
}

func MakesSimpleModifyResps() []core.OchanocoModifyResponse {
	simpleModifyResponse := func(proxy *core.OchanocoProxy, res *http.Response, c *gin.Context) bool {
		b := []byte(TEST_RESP_BODY2)
		res.Body = ioutil.NopCloser(bytes.NewReader(b))
		res.Header.Set("Content-Length", strconv.Itoa(len(b)))

		return core.FINISHED
	}

	return []core.OchanocoModifyResponse{simpleModifyResponse}
}

func MakeEmptyModifyResps() []core.OchanocoModifyResponse {
	simpleModifyResponse := func(proxy *core.OchanocoProxy, res *http.Response, c *gin.Context) bool {
		return core.FINISHED
	}

	return []core.OchanocoModifyResponse{simpleModifyResponse}
}
