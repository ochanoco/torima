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

func testServer(name, msg string, proxy *OchanocoProxy, t *testing.T) {
	t.Run(name, func(t *testing.T) {
		serv := httptest.NewServer(proxy.ReverseProxy)

		req, err := http.NewRequest(http.MethodGet, serv.URL, nil)
		if err != nil {
			t.Fatal(err)
		}

		resp, err := new(http.Client).Do(req)
		if err != nil {
			t.Fatal(err)
		}

		respBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Printf("msg buf: %v", respBytes)

		respStr := string(respBytes)

		fmt.Printf("msg buf: %v", respStr)

		if respStr != msg {
			t.Fatalf("wrong response: '%s'\nexpected: '%s'", respStr, TEST_RESP_BODY1)
		}
	})
}

func TestProxy(t *testing.T) {
	db, err := InitDB(TEST_DB_PATH)

	if err != nil {
		t.Fatalf("%v", err)
	}

	simpleServ := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		fmt.Fprint(writer, TEST_RESP_BODY1)
	}))

	simpleDirector := func(proxy *OchanocoProxy, req *http.Request) {
		url, err := url.Parse(simpleServ.URL)

		if err != nil {
			t.Fatalf("parse: %v", err)
		}

		fmt.Printf("scheme: %v\n", url.Scheme)

		req.URL.Scheme = url.Scheme
		req.URL.Host = url.Host
		req.URL.Path = "/"
	}

	simpleModifyResponse := func(proxy *OchanocoProxy, res *http.Response) {
		b := []byte(TEST_RESP_BODY2)
		res.Body = ioutil.NopCloser(bytes.NewReader(b))
		res.Header.Set("Content-Length", strconv.Itoa(len(b)))
	}

	directors := []OchanocoDirector{simpleDirector}
	modifyRespes := []OchanocoModifyResponse{}

	directorProxy := NewOchancoProxy(directors, modifyRespes, db)

	testServer("simple director", TEST_RESP_BODY1, &directorProxy, t)

	directors = []OchanocoDirector{simpleDirector}
	modifyRespes = []OchanocoModifyResponse{simpleModifyResponse}

	modifyProxy := NewOchancoProxy(directors, modifyRespes, db)

	testServer("simple director", TEST_RESP_BODY2, &modifyProxy, t)
}
