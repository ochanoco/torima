package core

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MainDirectorTester struct {
	Host       string
	Cookie     string
	ResultBody string
}

func (tester *MainDirectorTester) start(t *testing.T, proxy *OchanocoProxy, proxyServ *httptest.Server, testServ *httptest.Server) {
	testServUrl := ParseURL(t, testServ.URL)

	crp := proxy.Database.CreateServiceProvider(testServUrl.Host, testServUrl.Host)
	proxy.Database.SaveServiceProvider(crp)

	if tester.Host == "" {
		tester.Host = testServUrl.Host

	}

}
func (tester *MainDirectorTester) directors(t *testing.T, url string) []OchanocoDirector {
	return DEFAULT_DIRECTORS
}
func (tester *MainDirectorTester) modifyResps(t *testing.T) []OchanocoModifyResponse {
	return makeEmptyModifyResps()
}
func (tester *MainDirectorTester) testServers(t *testing.T) (*httptest.Server, *httptest.Server, *httptest.Server) {
	testServ := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		fmt.Fprint(writer, TEST_RESP_BODY1)

		userAgent := req.Header.Get("User-Agent")
		token := req.Header.Get("X-Ochanoco-Proxy-Token")

		if userAgent != "ochanoco" {
			t.Fatalf("wrong user agent: %s", userAgent)
		}

		if token != "<proxy_token>" {
			t.Fatalf("wrong proxy token: %s", token)
		}
	}))

	errorServ := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		fmt.Fprint(writer, TEST_RESP_ERROR_BODY)
	}))

	redirectServ := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		fmt.Fprint(writer, TEST_RESP_REDIRECT_BODY)
	}))

	return testServ, errorServ, redirectServ
}

func (tester *MainDirectorTester) request(t *testing.T, url string) *http.Response {
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Host = tester.Host

	req.Header.Set("Cookie", tester.Cookie)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		t.Fatal(err)
	}

	return resp
}

func (tester *MainDirectorTester) check(t *testing.T, resp *http.Response) {
	checkResponseWithBody(t, resp, tester.ResultBody)
}

func TestMainDirector(t *testing.T) {
	testerOk := MainDirectorTester{
		ResultBody: TEST_RESP_BODY1,
		Cookie:     "ochanoco-token=test",
	}
	runCommonTest(t, &testerOk, "main/director_ok")

	testerAuthFailed := MainDirectorTester{
		ResultBody: TEST_RESP_REDIRECT_BODY,
		Cookie:     "ochanoco-toke=",
	}
	runCommonTest(t, &testerAuthFailed, "main/director_auth_failed")

	testerSiteNotFound := MainDirectorTester{
		Host:       "www.example.com",
		ResultBody: TEST_RESP_ERROR_BODY,
		Cookie:     "ochanoco-token=test",
	}
	runCommonTest(t, &testerSiteNotFound, "main/director_site_not_found")

}
