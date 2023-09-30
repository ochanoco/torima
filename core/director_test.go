package core

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func directorSample(t *testing.T) (*http.Request, *gin.Context, *OchanocoProxy) {
	DB_TYPE = "sqlite3"
	DB_CONFIG = "../data/test.db?_fk=1"
	SECRET = "test_secret"

	proxy, err := ProxyServer()
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)

	req := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	return req, context, proxy
}

// test for RouteDirector
func TestRouteDirector(t *testing.T) {
	req, context, proxy := directorSample(t)
	c, err := RouteDirector("example.com", proxy, req, context)

	assert.NoError(t, err)
	assert.Equal(t, CONTINUE, c)
	assert.Equal(t, "example.com", req.URL.Host)
	assert.Equal(t, "https", req.URL.Scheme)
	assert.Equal(t, "ochanoco", req.Header.Get("User-Agent"))
	assert.Equal(t, SECRET, req.Header.Get("X-Ochanoco-Proxy-Token"))
}

// test for DefaultRouteDirector
func TestThirdPartyDirector(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	})

	ts := httptest.NewServer(h)
	defer ts.Close()

	u, err := url.Parse(ts.URL)
	assert.NoError(t, err)

	host := fmt.Sprintf("%v:%v", u.Host, u.Port())

	req, context, proxy := directorSample(t)
	req.URL.Path = "/ochanoco/redirect/" + host

	proxy.Config.ProtectionScope = []string{host}

	c, err := ThirdPartyDirector(proxy, req, context)
	assert.NoError(t, err)
	assert.Equal(t, CONTINUE, c)

	c, err = ThirdPartyDirector(proxy, req, context)
	assert.NoError(t, err)

	assert.Equal(t, CONTINUE, c)
	assert.Equal(t, host, req.URL.Host)
}

// test for DefaultRouteDirector
func TestThirdPartyDirectorNoParmit(t *testing.T) {
	unpermitHost := "not-in-list.example.com"

	req, context, proxy := directorSample(t)
	req.URL.Path = "/ochanoco/redirect/" + unpermitHost + "/"

	c, err := ThirdPartyDirector(proxy, req, context)
	assert.NoError(t, err)
	assert.Equal(t, CONTINUE, c)

	c, err = ThirdPartyDirector(proxy, req, context)
	assert.NoError(t, err)

	assert.Equal(t, CONTINUE, c)
	assert.NotEqual(t, unpermitHost, req.URL.Host)
}
