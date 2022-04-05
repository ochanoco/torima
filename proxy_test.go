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

type ProxyTester interface {
	setup(t *testing.T, servReturnBody *string)
	start(t *testing.T, proxyServ *httptest.Server, loginServ *httptest.Server, testServ *httptest.Server)
	check(t *testing.T, resp *http.Response)
	testServ(t *testing.T, writer http.ResponseWriter, req *http.Request)
}

func runTestProxyCommon(t *testing.T, tester ProxyTester, name string) {
	setupForTest()

	var servReturnBody = "ok"
	tester.setup(t, &servReturnBody)

	loginServ := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(writer, servReturnBody)
	}))

	testServ := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(writer, servReturnBody)
	}))

	rp := httputil.ReverseProxy{
		Director:       director,
		ModifyResponse: modifyResponse,
	}

	proxServ := httptest.NewServer(&rp)
	defer proxServ.Close()

	tester.start(t, proxServ, loginServ, testServ)

	LOGIN_REDIRECT_PAGE_URL = loginServ.URL

	t.Run(name, func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, proxServ.URL, nil)
		if err != nil {
			t.Error(err)
		}

		resp, err := new(http.Client).Do(req)
		if err != nil {
			t.Error(err)
		}

		tester.check(t, resp)
	})
}

/// Proxy Test (Result : OK)
type ProxyTestOKTester struct {
	respBody string
}

func NewProxyTestOKTester() ProxyTester {
	return &ProxyTestOKTester{""}
}

func (tester *ProxyTestOKTester) setup(t *testing.T, servReturnBody *string) {
	tester.respBody = "<p>ok</p>"

	*servReturnBody = tester.respBody
}

func (tester *ProxyTestOKTester) start(t *testing.T, proxyServ *httptest.Server, loginServ *httptest.Server, testServ *httptest.Server) {
	proxyDomain, err := url.Parse(proxyServ.URL)

	if err != nil {
		t.Errorf("failed parse %v", proxyServ.URL)
		return
	}

	testServDomain, err := url.Parse(testServ.URL)

	if err != nil {
		t.Errorf("failed parse %v", proxyServ.URL)
		return
	}

	proj := createProject(db, proxyDomain.Host, testServDomain.Host, "<line_id_for_proxy_ok_test>", "<name>")

	proj.Save(db.ctx)
}

func (tester *ProxyTestOKTester) check(t *testing.T, resp *http.Response) {
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	bs := string(b[:len(b)-1])

	if bs != tester.respBody {
		msg := fmt.Sprintf("wrong response: '%s'expected: '%s'", bs, tester.respBody)
		t.Error(msg)
	}
}

func (tester *ProxyTestOKTester) testServ(t *testing.T, writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(writer, tester.respBody)
	userAgent := req.Header.Get("User-Agent")
	token := req.Header.Get("X-BULLET-Proxy-Token")

	if userAgent != "bullet" {
		t.Errorf("wrong user agent: %s", userAgent)
	}

	if token != "<proxy_token>" {
		t.Errorf("wrong proxy token: %s", token)
	}
}

/// Proxy Test (Result : Fail Website Authentication)

type ProxyTestFailBecauseWebNotValid struct {
	testBody  string
	errorBody string
}

func NewProxyTestFailWebAuthTester() ProxyTester {
	return &ProxyTestFailBecauseWebNotValid{"", ""}
}

func (tester *ProxyTestFailBecauseWebNotValid) setup(t *testing.T, servReturnBody *string) {
	tester.testBody = "<p>ok</p>"
	tester.errorBody = "<p>error</p>"

	*servReturnBody = tester.testBody

	errorServ := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(writer, tester.errorBody)
	}))

	ERROR_PAGE_URL = errorServ.URL
}

func (tester *ProxyTestFailBecauseWebNotValid) start(t *testing.T, proxyServ *httptest.Server, loginServ *httptest.Server, testServ *httptest.Server) {
}

func (tester *ProxyTestFailBecauseWebNotValid) check(t *testing.T, resp *http.Response) {
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	t.Errorf("hello: %v", b)

	bs := string(b[:len(b)-1])

	if bs != tester.errorBody {
		msg := fmt.Sprintf("wrong response: '%s'\nexpected: '%s'", bs, tester.errorBody)
		t.Error(msg)
	}
}

func (tester *ProxyTestFailBecauseWebNotValid) testServ(t *testing.T, writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(writer, tester.testBody)
}

// proxyDomain, err := url.Parse(proxServ.URL)

// if err != nil {
// 	t.Errorf("failed parse %v", proxServ.URL)
// 	return
// }
// 	proj := createProject(db, proxyDomain.Host, "test_for_proxy_ng1_test.example.com", "<line_id_for_proxy_ng1_test>", "<name>")

// proj.Save(db.ctx)

// if err != nil {
// 	t.Errorf("failed creating white list: %v", err)
// 	return
// }

// type ProxyTestOKTester struct{}

// func (p *ProxyTestOKTester) setup(param ProxyTesterParam) {}
// func (p *ProxyTestOKTester) check(param ProxyTesterParam, resp *http.Response) {
// 	expectedBody := ""

// 	b, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		param.t.Error(err)
// 	}
// 	bs := string(b[:len(b)-1])

// 	if bs != expectedBody {
// 		param.t.Errorf("wrong response: '%s'\nexpected: '%s'", bs, expectedBody)
// 	}
// }

// func TestProxyRedirectToLogin(t *testing.T) {
// 	setupForTest()

// 	loginBody := "<body>login</body>"

// 	loginServ := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
// 		fmt.Fprintln(writer, loginBody)
// 	}))

// 	t.Run("test proxy", func(t *testing.T) {
// 		LOGIN_REDIRECT_PAGE_URL = loginServ.URL

// 		rp := httputil.ReverseProxy{
// 			Director:       director,
// 			ModifyResponse: modifyResponse,
// 		}

// 		serv := httptest.NewServer(&rp)

// 		defer serv.Close()

// 		proxyDomain, err := url.Parse(serv.URL)

// 		if err != nil {
// 			t.Errorf("failed parse %v", serv.URL)
// 			return
// 		}

// 		proj := createProject(db, proxyDomain.Host, "test_for_proxy_ng1_test.example.com", "<line_id_for_proxy_ng1_test>", "<name>")

// 		proj.Save(db.ctx)

// 		if err != nil {
// 			t.Errorf("failed creating white list: %v", err)
// 			return
// 		}

// 		req, err := http.NewRequest(http.MethodPost, serv.URL, nil)
// 		if err != nil {
// 			t.Error(err)
// 		}

// 		resp, err := new(http.Client).Do(req)
// 		if err != nil {
// 			t.Error(err)
// 		}

// 		b, err := ioutil.ReadAll(resp.Body)
// 		if err != nil {
// 			t.Error(err)
// 		}

// 		bs := string(b[:len(b)-1])

// 		if bs != loginBody {
// 			msg := fmt.Sprintf("wrong response: '%s'\nexpected: '%s'", bs, loginBody)
// 			t.Error(msg)
// 		}
// 	})
// }

func TestProxyFailLogin(t *testing.T) {
	setupForTest()

	errorBody := "<body>error</body>"

	errorServ := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(writer, errorBody)
	}))

	ERROR_PAGE_URL = errorServ.URL

	t.Run("test proxy", func(t *testing.T) {
		ERROR_PAGE_URL = errorServ.URL

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

func TestProxy(t *testing.T) {
	// var tester_website_not_valid = NewProxyTestFailWebAuthTester()
	// runTestProxyCommon(t, tester_website_not_valid, "test_website_not_valid")

	// var tester_ok = NewProxyTestOKTester()
	// runTestProxyCommon(t, tester_ok, "test_proxy_ok")
}
