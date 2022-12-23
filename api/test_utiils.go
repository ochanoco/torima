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
	setup(t *testing.T)
	start(t *testing.T, proxyServ *httptest.Server, testServ *httptest.Server)

	director(t *testing.T, URL string) OchanocoDirector
	modifyResp(t *testing.T) OchanocoModifyResponse
	testServ(t *testing.T) *httptest.Server

	request(t *testing.T, url string) *http.Response
	check(t *testing.T, resp *http.Response)
}

func runCommonTest(t *testing.T, tester proxyTester, name string) {
	db, err := InitDB(TEST_DB_PATH)

	if err != nil {
		t.Fatalf("%v", err)
	}

	tester.setup(t)

	testServ := tester.testServ(t)
	director := tester.director(t, testServ.URL)

	modifyResp := tester.modifyResp(t)

	directors := []OchanocoDirector{director}
	modifyRespes := []OchanocoModifyResponse{modifyResp}

	proxy := NewOchancoProxy(directors, modifyRespes, db)

	t.Run(name, func(t *testing.T) {
		serv := httptest.NewServer(proxy.ReverseProxy)
		resp := tester.request(t, serv.URL)

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

func makeCheckResponseWithBody(t *testing.T, resp *http.Response, expect string) func(t *testing.T, resp *http.Response) {
	f := func(t *testing.T, resp *http.Response) {
		respBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("%v", err)
		}

		respStr := string(respBytes)

		if respStr != expect {
			t.Fatalf("wrong response: '%s'expected: '%s'", respStr, expect)
		}
	}

	return f
}

func makeSimpleServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		fmt.Fprint(writer, TEST_RESP_BODY1)
	}))

}

func makesSimpleDirector(t *testing.T, URL string) OchanocoDirector {
	test := func(proxy *OchanocoProxy, req *http.Request) {
		url, err := url.Parse(URL)

		if err != nil {
			t.Fatalf("parse: %v", err)
		}

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
