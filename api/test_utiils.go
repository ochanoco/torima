package main

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
)

const TEST_RESP_BODY1 = "hoge"
const TEST_RESP_BODY2 = "piyo"
const TEST_RESP_ERROR_BODY = "error"
const TEST_RESP_REDIRECT_BODY = "redirect"

type proxyTester interface {
	start(t *testing.T, proxy *OchanocoProxy, proxyServ *httptest.Server, testServ *httptest.Server)

	directors(t *testing.T, URL string) []OchanocoDirector
	modifyResps(t *testing.T) []OchanocoModifyResponse
	testServers(t *testing.T) (*httptest.Server, *httptest.Server, *httptest.Server)

	request(t *testing.T, url string) *http.Response
	check(t *testing.T, resp *http.Response)
}

func runCommonTest(t *testing.T, tester proxyTester, name string) {
	db, err := InitDB(DB_CONFIG)

	if err != nil {
		t.Fatalf("%v", err)
	}

	testServ, errorServ, redirectServ := tester.testServers(t)
	ERROR_PAGE_URL = errorServ.URL
	LOGIN_REDIRECT_PAGE_URL = redirectServ.URL

	directors := tester.directors(t, testServ.URL)

	modifyResps := tester.modifyResps(t)

	r := gin.Default()

	proxy := NewOchancoProxy(r, directors, modifyResps, db)
	proxyServ := httptest.NewServer(proxy.Engine)

	tester.start(t, &proxy, proxyServ, testServ)

	t.Run(name, func(t *testing.T) {
		resp := tester.request(t, proxyServ.URL)

		tester.check(t, resp)
	})
}

/////////

func requestGetforTest(t *testing.T, url string) *http.Response {
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

func requestPostWithCookieForTest(t *testing.T, url, cookie string) *http.Response {
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

func checkResponseWithBody(t *testing.T, resp *http.Response, expect string) {
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("%v", err)
	}

	respStr := string(respBytes)

	if respStr != expect {
		t.Fatalf("wrong response: '%s'expected: '%s'", respStr, expect)
	}
}

func makeSimpleServers() (*httptest.Server, *httptest.Server, *httptest.Server) {
	s := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		fmt.Fprint(writer, TEST_RESP_BODY1)
	}))

	return s, s, s
}

func makesSimpleDirectors(t *testing.T, URL string) []OchanocoDirector {
	test := func(proxy *OchanocoProxy, req *http.Request, c *gin.Context) bool {
		url := parseURL(t, URL)

		req.URL.Scheme = url.Scheme
		req.URL.Host = url.Host
		req.URL.Path = "/"

		return FINISHED
	}

	return []OchanocoDirector{test}
}

func makesSimpleModifyResps() []OchanocoModifyResponse {
	simpleModifyResponse := func(proxy *OchanocoProxy, res *http.Response, c *gin.Context) bool {
		b := []byte(TEST_RESP_BODY2)
		res.Body = ioutil.NopCloser(bytes.NewReader(b))
		res.Header.Set("Content-Length", strconv.Itoa(len(b)))

		return FINISHED
	}

	return []OchanocoModifyResponse{simpleModifyResponse}
}

func makeEmptyModifyResps() []OchanocoModifyResponse {
	simpleModifyResponse := func(proxy *OchanocoProxy, res *http.Response, c *gin.Context) bool {
		return FINISHED
	}

	return []OchanocoModifyResponse{simpleModifyResponse}
}

func parseURL(t *testing.T, URL string) *url.URL {
	url, err := url.Parse(URL)

	if err != nil {
		t.Fatalf("parse: %v", err)
	}

	return url
}
