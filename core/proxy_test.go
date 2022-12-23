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
)

func TestProxyDirectorAndModifyResponse(t *testing.T) {
	db, err := InitDB(TEST_DB_PATH)

	if err != nil {
		t.Fatalf("%v", err)
	}

	directors := []OchanocoDirector{}
	modifyRespes := []OchanocoResponse{}

	msg := "hello"

	simpleServ := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(writer, msg)
	}))

	simpleDirector := func(proxy *OchanocoProxy, req *http.Request) {
		url, _ := url.Parse(simpleServ.URL)
		req.URL.Scheme = url.Scheme
		req.URL.Host = url.Host
		req.URL.Path = "/"
	}

	simpleModifyResponse := func(proxy *OchanocoProxy, res *http.Response) {
		b := []byte(msg)
		res.Body = ioutil.NopCloser(bytes.NewReader(b))
		res.Header.Set("Content-Length", strconv.Itoa(len(b)))
	}

	proxy := NewOchancoProxy(directors, modifyRespes, db)
	proxy.AddDirector(simpleDirector)

	targetServ := httptest.NewServer(proxy.ReverseProxy)

	t.Run("simple director", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, targetServ.URL, nil)
		if err != nil {
			t.Error(err)
		}

		resp, err := new(http.Client).Do(req)
		if err != nil {
			t.Error(err)
		}

		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}

		bs := string(b[:len(b)-1])

		if bs != msg {
			msg := fmt.Sprintf("wrong response: '%s'\nexpected: '%s'", bs, msg)
			t.Error(msg)
		}
	})

	proxy.AddModifyResponse(simpleModifyResponse)
	t.Run("modify response", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, targetServ.URL, nil)
		if err != nil {
			t.Error(err)
		}

		resp, err := new(http.Client).Do(req)
		if err != nil {
			t.Error(err)
		}

		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}

		bs := string(b)

		if bs != msg {
			msg := fmt.Sprintf("wrong response: '%s'\nexpected: '%s'", bs, msg)
			t.Error(msg)
		}
	})
}
