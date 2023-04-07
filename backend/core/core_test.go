package core

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

////////// CoreDirectorTester

type CoreDirectorTester struct{}

func (tester *CoreDirectorTester) start(t *testing.T, proxy *OchanocoProxy, proxyServ *httptest.Server, testServ *httptest.Server) {
}

func (tester *CoreDirectorTester) directors(t *testing.T, url string) []OchanocoDirector {
	return makesSimpleDirectors(t, url)
}
func (tester *CoreDirectorTester) modifyResps(t *testing.T) []OchanocoModifyResponse {
	return makeEmptyModifyResps()
}
func (tester *CoreDirectorTester) testServers(t *testing.T) (*httptest.Server, *httptest.Server, *httptest.Server) {
	return makeSimpleServers()
}

func (tester *CoreDirectorTester) request(t *testing.T, url string) *http.Response {
	return requestGetforTest(t, url)
}
func (tester *CoreDirectorTester) check(t *testing.T, resp *http.Response) {
	checkResponseWithBody(t, resp, TEST_RESP_BODY1)
}

func TestCoreDirector(t *testing.T) {
	tester := CoreDirectorTester{}
	runCommonTest(t, &tester, "core/director")
}

// //////// CoreModifyResponseTester
type CoreModifyResponseTester struct{}

func (tester *CoreModifyResponseTester) start(t *testing.T, proxy *OchanocoProxy, proxyServ *httptest.Server, testServ *httptest.Server) {
}

func (tester *CoreModifyResponseTester) directors(t *testing.T, url string) []OchanocoDirector {
	return makesSimpleDirectors(t, url)
}
func (tester *CoreModifyResponseTester) modifyResps(t *testing.T) []OchanocoModifyResponse {
	return makesSimpleModifyResps()
}
func (tester *CoreModifyResponseTester) testServers(t *testing.T) (*httptest.Server, *httptest.Server, *httptest.Server) {
	return makeSimpleServers()
}

func (tester *CoreModifyResponseTester) request(t *testing.T, url string) *http.Response {
	return requestGetforTest(t, url)
}
func (tester *CoreModifyResponseTester) check(t *testing.T, resp *http.Response) {
	checkResponseWithBody(t, resp, TEST_RESP_BODY2)
}

func TestCoreModifyResp(t *testing.T) {
	tester := CoreModifyResponseTester{}
	runCommonTest(t, &tester, "core/modify_resp")
}
