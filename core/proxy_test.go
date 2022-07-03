package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
	"strconv"
	"testing"
)

func TestProxyDirectorAndModifyResponse(t *testing.T) {
	Directors = []func(req *http.Request){}
	ModifyResponses = []func(req *http.Response){}

	msg := "hello"

	simpleServ := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(writer, msg)
	}))

	simpleDirector := func(req *http.Request) {
		url, _ := url.Parse(simpleServ.URL)
		req.URL.Scheme = url.Scheme
		req.URL.Host = url.Host
		req.URL.Path = "/"
	}

	simpleModifyResponse := func(res *http.Response) {
		b := []byte(msg)
		res.Body = ioutil.NopCloser(bytes.NewReader(b))
		res.Header.Set("Content-Length", strconv.Itoa(len(b)))
	}

	rp := httputil.ReverseProxy{
		Director:       director,
		ModifyResponse: modifyResponse,
	}

	Directors = append(Directors, simpleDirector)
	targetServ := httptest.NewServer(&rp)

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

	ModifyResponses = append(ModifyResponses, simpleModifyResponse)

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
