package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"testing"
)

func TestProxy(t *testing.T) {
	expected := "<html><head></head><body>hello</body></html>"

	t.Run("test proxy", func(t *testing.T) {
		testServ := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
			fmt.Fprintln(writer, "<body>hi</body>")
			userAgent := req.Header.Get("User-Agent")
			userId := req.Header.Get("X-TUASET-User-ID")
			token := req.Header.Get("X-TUASET-Proxy-Token")

			if userAgent != "tuaset" {
				msg := fmt.Sprintf("wrong user agent: %s", userAgent)
				t.Error(msg)
			}

			if userId != "<user_id>" {
				msg := fmt.Sprintf("wrong user agent: %s", userId)
				t.Error(msg)
			}

			if token != "<token>" {
				msg := fmt.Sprintf("wrong user agent: %s", token)
				t.Error(msg)
			}
		}))

		LOGIN_PAGE_URL = testServ.URL

		rp := httputil.ReverseProxy{
			Director:       director,
			ModifyResponse: modifyResponse,
		}

		serv := httptest.NewServer(&rp)

		defer serv.Close()

		req, err := http.NewRequest(http.MethodGet, serv.URL, nil)
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

		if bs != expected {
			msg := fmt.Sprintf("wrong response: %s", bs)
			t.Error(msg)
		}
	})
}
