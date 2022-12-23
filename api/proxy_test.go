package proxy

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MainDirectorTesterOk struct {
	host string
}

func (tester *MainDirectorTesterOk) start(t *testing.T, proxy *OchanocoProxy, proxyServ *httptest.Server, testServ *httptest.Server) {
	testServUrl := parseURL(t, testServ.URL)

	crp := proxy.Database.CreateServiceProvider(testServUrl.Host, testServUrl.Host)
	proxy.Database.SaveServiceProvider(crp)

	tester.host = testServUrl.Host

}
func (tester *MainDirectorTesterOk) director(t *testing.T, url string) OchanocoDirector {
	return MainDirector
}
func (tester *MainDirectorTesterOk) modifyResp(t *testing.T) OchanocoModifyResponse {
	return makeEmptyModifyResp()
}
func (tester *MainDirectorTesterOk) testServ(t *testing.T) *httptest.Server {
	f := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(writer, TEST_RESP_BODY1)

		userAgent := req.Header.Get("User-Agent")
		token := req.Header.Get("X-Ochanoco-Proxy-Token")

		if userAgent != "ochanoco" {
			t.Fatalf("wrong user agent: %s", userAgent)
		}

		if token != "<proxy_token>" {
			t.Fatalf("wrong proxy token: %s", token)
		}
	}))

	return f
}

func (tester *MainDirectorTesterOk) request(t *testing.T, url string) *http.Response {
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Host = tester.host

	req.Header.Set("Cookie", "ochanoco-token=test")

	resp, err := new(http.Client).Do(req)
	if err != nil {
		t.Fatal(err)
	}

	return resp
}

func (tester *MainDirectorTesterOk) check(t *testing.T, resp *http.Response) {
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status code not match: %v", resp.StatusCode)
	}
}

func TestMainDirector(t *testing.T) {
	tester := MainDirectorTesterOk{}
	runCommonTest(t, &tester, "main/director")
}

// func startServers(t *testing.T) {
// 	db, err := InitDB(TEST_DB_PATH)
// 	if err != nil {
// 		t.Fatalf("fail to initialize: %v", err)
// 	}

// 	testServ := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
// 		tester.testServ(t, writer, req)
// 	}))

// 	ochanocoProxy := NewOchancoProxy(
// 		[]OchanocoDirector{MainDirector},
// 		[]OchanocoModifyResponse{},
// 		db,
// 	)

// 	proxyDomain, err := url.Parse(proxyServ.URL)

// 	if err != nil {
// 		t.Errorf("failed parse %v", proxyServ.URL)
// 		return
// 	}

// 	testServDomain, err := url.Parse(testServ.URL)

// 	if err != nil {
// 		t.Errorf("failed parse %v", testServ.URL)
// 		return
// 	}
// }

// func runTest(t *testing.T, tester ProxyTester, name string) {
// 	proxServ := httptest.NewServer(ochanocoProxy.ReverseProxy)
// 	defer proxServ.Close()

// 	tester.start(t, proxServ, testServ)

// 	t.Run(name, func(t *testing.T) {
// 		req, err := http.NewRequest(http.MethodPost, proxServ.URL, nil)
// 		if err != nil {
// 			t.Error(err)
// 		}

// 		req.Header.Set("Cookie", cookie)

// 		resp, err := new(http.Client).Do(req)
// 		if err != nil {
// 			t.Error(err)
// 		}

// 		tester.check(t, resp)
// 	})
// }

// type ProxyTestOKTester struct {
// 	respBody string
// 	lineId   string
// 	name     string
// }

// func NewProxyTestOKTester() ProxyTester {
// 	return &ProxyTestOKTester{"<p>ok</p>", "test_line_id_for_test_ok", "test_name_for_test_ok"}
// }

// func (tester *ProxyTestOKTester) setup(t *testing.T, servReturnBody *string, cookie *string) {
// 	*servReturnBody = tester.respBody
// 	*cookie = fmt.Sprintf("ochanoco-token=%s", tester.lineId)
// }

// func (tester *ProxyTestOKTester) start(t *testing.T, proxyServ *httptest.Server, testServ *httptest.Server) {
// 	proxyDomain, err := url.Parse(proxyServ.URL)

// 	if err != nil {
// 		t.Errorf("failed parse %v", proxyServ.URL)
// 		return
// 	}

// 	testServDomain, err := url.Parse(testServ.URL)

// 	if err != nil {
// 		t.Errorf("failed parse %v", testServ.URL)
// 		return
// 	}

// 	proj := db.CreateServiceProvider(proxyDomain.Host, testServDomain.Host)

// 	_, err = proj.Save(db.ctx)

// 	if err != nil {
// 		t.Errorf("failed save %v", err)
// 		return
// 	}
// }

// // // / Proxy Test (Result : Fail Website Authentication)
// // type ProxyTestFailBecauseWebNotValid struct {
// // 	testBody  string
// // 	errorBody string
// // }

// // func NewProxyTestFailWebAuthTester() ProxyTester {
// // 	return &ProxyTestFailBecauseWebNotValid{"", ""}
// // }

// // func (tester *ProxyTestFailBecauseWebNotValid) setup(t *testing.T, servReturnBody *string, cookie *string) {
// // 	tester.testBody = "<p>ok</p>"
// // 	tester.errorBody = "<p>error</p>"

// // 	*servReturnBody = tester.testBody

// // 	errorServ := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
// // 		fmt.Fprintln(writer, tester.errorBody)
// // 	}))

// // 	ERROR_PAGE_URL = errorServ.URL
// // }

// // func (tester *ProxyTestFailBecauseWebNotValid) start(t *testing.T, proxyServ *httptest.Server, testServ *httptest.Server) {
// // }

// // func (tester *ProxyTestFailBecauseWebNotValid) check(t *testing.T, resp *http.Response) {
// // 	b, err := ioutil.ReadAll(resp.Body)
// // 	if err != nil {
// // 		t.Error(err)
// // 	}

// // 	bs := string(b[:len(b)-1])

// // 	if bs != tester.errorBody {
// // 		msg := fmt.Sprintf("wrong response: '%s'\nexpected: '%s'", bs, tester.errorBody)
// // 		t.Error(msg)
// // 	}
// // }

// // func (tester *ProxyTestFailBecauseWebNotValid) testServ(t *testing.T, writer http.ResponseWriter, req *http.Request) {
// // 	fmt.Fprintln(writer, tester.testBody)
// // }

// // // / Proxy Test (Result : Redicrect To Login Page)
// // type ProxyTestRedirectLogin struct {
// // 	testBody  string
// // 	errorBody string
// // }

// // func NewProxyTestRedirectLogin() ProxyTester {
// // 	return &ProxyTestFailBecauseWebNotValid{"", ""}
// // }

// // func (tester *ProxyTestRedirectLogin) setup(t *testing.T, servReturnBody *string) {
// // 	tester.testBody = "<p>ok</p>"
// // 	tester.errorBody = "<p>login</p>"

// // 	*servReturnBody = tester.testBody

// // 	loginServ := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
// // 		fmt.Fprintln(writer, tester.errorBody)
// // 	}))

// // 	LOGIN_REDIRECT_PAGE_URL = loginServ.URL
// // }

// // func (tester *ProxyTestRedirectLogin) start(t *testing.T, proxyServ *httptest.Server, loginServ *httptest.Server, testServ *httptest.Server) {
// // }

// // func (tester *ProxyTestRedirectLogin) check(t *testing.T, resp *http.Response) {
// // 	b, err := ioutil.ReadAll(resp.Body)
// // 	if err != nil {
// // 		t.Error(err)
// // 	}

// // 	bs := string(b[:len(b)-1])

// // 	if bs != tester.errorBody {
// // 		msg := fmt.Sprintf("wrong response: '%s'\nexpected: '%s'", bs, tester.errorBody)
// // 		t.Error(msg)
// // 	}
// // }
// //
// // func (tester *ProxyTestRedirectLogin) testServ(t *testing.T, writer http.ResponseWriter, req *http.Request) {
// // 	fmt.Fprintln(writer, tester.testBody)
// // }

// func TestProxies(t *testing.T) {
// 	// d, err := InitDB(TEST_DB_PATH)
// 	// if err != nil {
// 	// 	t.Fatalf("failed to init db: %v", err)
// 	// }

// 	// db = d

// 	// var tester_website_not_valid = NewProxyTestFailWebAuthTester()
// 	// runTestProxyCommon(t, tester_website_not_valid, "test_website_not_valid")

// 	// var tester_redirect_to_login = NewProxyTestFailWebAuthTester()
// 	// runTestProxyCommon(t, tester_redirect_to_login, "tester_redirect_to_login")

// 	// var tester_ok = NewProxyTestOKTester()
// 	// runTestProxyCommon(t, tester_ok, "test_proxy_ok")
// }
