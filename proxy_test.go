package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"testing"
)

func TestProxyAuthOK(t *testing.T) {
	normalBody := "<body>ok</body>"

	t.Run("test proxy", func(t *testing.T) {
		testServ := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
			fmt.Fprintln(writer, normalBody)
			userAgent := req.Header.Get("User-Agent")
			token := req.Header.Get("X-BULLET-Proxy-Token")

			if userAgent != "bullet" {
				msg := fmt.Sprintf("wrong user agent: %s", userAgent)
				t.Error(msg)
			}

			if token != "<proxy_token>" {
				msg := fmt.Sprintf("wrong proxy token: %s", token)
				t.Error(msg)
			}
		}))

		EXAMPLE_URL = testServ.URL

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

		bs := string(b[:len(b)-1])

		if bs != normalBody {
			msg := fmt.Sprintf("wrong response: '%s'expected: '%s'", bs, normalBody)
			t.Error(msg)
		}
	})
}

func TestProxyRedirectToLogin(t *testing.T) {
	loginBody := "<body>login</body>"

	loginServ := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(writer, loginBody)
	}))

	t.Run("test proxy", func(t *testing.T) {
		LOGIN_REDIRECT_PAGE_URL = loginServ.URL

		rp := httputil.ReverseProxy{
			Director:       director,
			ModifyResponse: modifyResponse,
		}

		serv := httptest.NewServer(&rp)

		defer serv.Close()

		req, err := http.NewRequest(http.MethodPost, serv.URL, nil)
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

		if bs != loginBody {
			msg := fmt.Sprintf("wrong response: '%s'\nexpected: '%s'", bs, loginBody)
			t.Error(msg)
		}
	})
}
