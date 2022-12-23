package proxy

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
)

const TEST_RESP_BODY1 = "hoge"
const TEST_RESP_BODY2 = "piyo"

type proxyTester interface {
	start(t *testing.T, proxy *OchanocoProxy, proxyServ *httptest.Server, testServ *httptest.Server)

	director(t *testing.T, URL string) OchanocoDirector
	modifyResp(t *testing.T) OchanocoModifyResponse
	testServ(t *testing.T) *httptest.Server

	request(t *testing.T, url string) *http.Response
	check(t *testing.T, resp *http.Response)
}

func runCommonTest(t *testing.T, tester proxyTester, name string) {
	db, err := InitDB(DB_CONFIG)

	if err != nil {
		t.Fatalf("%v", err)
	}

	testServ := tester.testServ(t)
	director := tester.director(t, testServ.URL)

	modifyResp := tester.modifyResp(t)

	directors := []OchanocoDirector{director}
	modifyRespes := []OchanocoModifyResponse{modifyResp}

	proxy := NewOchancoProxy(directors, modifyRespes, db)
	proxyServ := httptest.NewServer(proxy.ReverseProxy)

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

func makeSimpleServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		fmt.Fprint(writer, TEST_RESP_BODY1)
	}))

}

func makesSimpleDirector(t *testing.T, URL string) OchanocoDirector {
	test := func(proxy *OchanocoProxy, req *http.Request) {
		url := parseURL(t, URL)

		req.URL.Scheme = url.Scheme
		req.URL.Host = url.Host
		req.URL.Path = "/"
	}

	return test
}

func makesSimpleModifyResp() OchanocoModifyResponse {
	simpleModifyResponse := func(proxy *OchanocoProxy, res *http.Response) {
		b := []byte(TEST_RESP_BODY2)
		res.Body = ioutil.NopCloser(bytes.NewReader(b))
		res.Header.Set("Content-Length", strconv.Itoa(len(b)))
	}

	return simpleModifyResponse
}

func makeEmptyModifyResp() OchanocoModifyResponse {
	simpleModifyResponse := func(proxy *OchanocoProxy, res *http.Response) {
	}

	return simpleModifyResponse
}

func parseURL(t *testing.T, URL string) *url.URL {
	url, err := url.Parse(URL)

	if err != nil {
		t.Fatalf("parse: %v", err)
	}

	return url
}
