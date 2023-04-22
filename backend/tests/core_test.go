package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ochanoco/proxy/core"
)

////////// CoreDirectorTester

type CoreDirectorTester struct{}

func (tester *CoreDirectorTester) Start(t *testing.T, proxy *core.OchanocoProxy, proxyServ *httptest.Server, testServ *httptest.Server) {
}

func (tester *CoreDirectorTester) Directors(t *testing.T, url string) []core.OchanocoDirector {
	return MakesSimpleDirectors(t, url)
}
func (tester *CoreDirectorTester) ModifyResps(t *testing.T) []core.OchanocoModifyResponse {
	return MakeEmptyModifyResps()
}
func (tester *CoreDirectorTester) TestServers(t *testing.T) (*httptest.Server, *httptest.Server, *httptest.Server) {
	return MakeSimpleServers()
}

func (tester *CoreDirectorTester) Request(t *testing.T, url string) *http.Response {
	return RequestGetforTest(t, url)
}
func (tester *CoreDirectorTester) Check(t *testing.T, resp *http.Response) {
	CheckResponseWithBody(t, resp, TEST_RESP_BODY1)
}

func TestCoreDirector(t *testing.T) {
	tester := CoreDirectorTester{}
	RunCommonTest(t, &tester, "core/director")
}

// //////// CoreModifyResponseTester
type CoreModifyResponseTester struct{}

func (tester *CoreModifyResponseTester) Start(t *testing.T, proxy *core.OchanocoProxy, proxyServ *httptest.Server, testServ *httptest.Server) {
}

func (tester *CoreModifyResponseTester) Directors(t *testing.T, url string) []core.OchanocoDirector {
	return MakesSimpleDirectors(t, url)
}
func (tester *CoreModifyResponseTester) ModifyResps(t *testing.T) []core.OchanocoModifyResponse {
	return MakesSimpleModifyResps()
}
func (tester *CoreModifyResponseTester) TestServers(t *testing.T) (*httptest.Server, *httptest.Server, *httptest.Server) {
	return MakeSimpleServers()
}

func (tester *CoreModifyResponseTester) Request(t *testing.T, url string) *http.Response {
	return RequestGetforTest(t, url)
}
func (tester *CoreModifyResponseTester) Check(t *testing.T, resp *http.Response) {
	CheckResponseWithBody(t, resp, TEST_RESP_BODY2)
}

func TestCoreModifyResp(t *testing.T) {
	tester := CoreModifyResponseTester{}
	RunCommonTest(t, &tester, "core/modify_resp")
}
