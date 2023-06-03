package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ochanoco/proxy/core"
)

type MainDirectorTester struct {
	Host       string
	Cookie     string
	ResultBody string
}

func (tester *MainDirectorTester) Start(t *testing.T, proxy *core.OchanocoProxy, proxyServ *httptest.Server, testServ *httptest.Server) {
	testServUrl := ParseURL(t, testServ.URL)

	crp := proxy.Database.CreateServiceProvider(testServUrl.Host, testServUrl.Host)
	proxy.Database.SaveServiceProvider(crp)

	if tester.Host == "" {
		tester.Host = testServUrl.Host

	}
}
func (tester *MainDirectorTester) Directors(t *testing.T, url string) []core.OchanocoDirector {
	return core.DEFAULT_DIRECTORS
}
func (tester *MainDirectorTester) ModifyResps(t *testing.T) []core.OchanocoModifyResponse {
	return MakeEmptyModifyResps()
}
func (tester *MainDirectorTester) TestServers(t *testing.T) (*httptest.Server, *httptest.Server, *httptest.Server) {
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

func (tester *MainDirectorTester) Request(t *testing.T, url string) *http.Response {
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

func (tester *MainDirectorTester) Check(t *testing.T, resp *http.Response) {
	CheckResponseWithBody(t, resp, tester.ResultBody)
}
