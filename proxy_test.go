package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
	"testing"
)

func TestProxyAuthOK(t *testing.T) {
	setupForTest()

	normalBody := "<body>ok</body>"

	t.Run("test proxy", func(t *testing.T) {
		testServ := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
			fmt.Fprintln(writer, normalBody)
			userAgent := req.Header.Get("User-Agent")
			token := req.Header.Get("X-BULLET-Proxy-Token")

			if userAgent != "bullet" {
				t.Errorf("wrong user agent: %s", userAgent)
			}

			if token != "<proxy_token>" {
				t.Errorf("wrong proxy token: %s", token)
			}
		}))

		rp := httputil.ReverseProxy{
			Director:       director,
			ModifyResponse: modifyResponse,
		}

		serv := httptest.NewServer(&rp)

		defer serv.Close()

		proxyDomain, err := url.Parse(serv.URL)

		if err != nil {
			t.Errorf("failed parse %v", serv.URL)
			return
		}

		testServDomain, err := url.Parse(testServ.URL)

		if err != nil {
			t.Errorf("failed parse: %v (%v)", testServ.URL, err)
			return
		}

		proj := createProject(db, proxyDomain.Host, testServDomain.Host, "<line_id_for_proxy_ok_test>", "<name>")

		proj.Save(db.ctx)

		if err != nil {
			t.Errorf("failed creating white list: %v", err)
			return
		}

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
	setupForTest()

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

		proxyDomain, err := url.Parse(serv.URL)

		if err != nil {
			t.Errorf("failed parse %v", serv.URL)
			return
		}

		proj := createProject(db, proxyDomain.Host, "test_for_proxy_ng1_test.example.com", "<line_id_for_proxy_ng1_test>", "<name>")

		proj.Save(db.ctx)

		if err != nil {
			t.Errorf("failed creating white list: %v", err)
			return
		}

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

func TestProxyFailLogin(t *testing.T) {
	setupForTest()

	errorBody := "<body>error</body>"

	errorServ := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(writer, errorBody)
	}))

	ERROR_PAGE_URL = errorServ.URL

	t.Run("test proxy", func(t *testing.T) {
		LOGIN_REDIRECT_PAGE_URL = errorServ.URL

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

		if bs != errorBody {
			msg := fmt.Sprintf("wrong response: '%s'\nexpected: '%s'", bs, errorBody)
			t.Error(msg)
		}
	})
}
