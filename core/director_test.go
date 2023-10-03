package core

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func directorSample(t *testing.T) (*http.Request, *TestResponseRecorder, *gin.Context, *OchanocoProxy) {
	DB_TYPE = "sqlite3"
	DB_CONFIG = "../data/test.db?_fk=1"
	SECRET = "test_secret"

	recorder := CreateTestResponseRecorder()
	context, r := gin.CreateTestContext(recorder)

	store := cookie.NewStore([]byte("test"))
	r.Use(sessions.Sessions("ochanoco-session", store))

	db, err := InitDB(DB_CONFIG)
	assert.NoError(t, err)

	config, file, err := readTestConfig(t)
	assert.NoError(t, err)
	defer os.Remove(file.Name())

	proxy := NewOchancoProxy(r, DEFAULT_DIRECTORS, DEFAULT_MODIFY_RESPONSES, DEFAULT_PROXYWEB_PAGES, config, db)
	req := httptest.NewRequest("GET", "http://localhost:8080/", nil)

	return req, recorder, context, &proxy
}

func setupMockServer(handler http.HandlerFunc, req *http.Request, t *testing.T) (*httptest.Server, *url.URL) {
	h := http.HandlerFunc(handler)

	ts := httptest.NewServer(h)
	u, err := url.Parse(ts.URL)
	assert.NoError(t, err)

	req.URL.Path = "/hello"
	req.URL.Host = u.Host
	req.Host = u.Host

	return ts, u
}

// test for RouteDirector
func TestRouteDirector(t *testing.T) {
	req, _, context, proxy := directorSample(t)
	c, err := RouteDirector("example.com", proxy, req, context)

	assert.NoError(t, err)
	assert.Equal(t, CONTINUE, c)
	assert.Equal(t, "example.com", req.URL.Host)
	assert.Equal(t, "http", req.URL.Scheme)
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

	req, _, context, proxy := directorSample(t)

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

	req, _, context, proxy := directorSample(t)

	req.URL.Path = "/ochanoco/redirect/" + unpermitHost + "/"

	c, err := ThirdPartyDirector(proxy, req, context)
	assert.NoError(t, err)
	assert.Equal(t, CONTINUE, c)

	c, err = ThirdPartyDirector(proxy, req, context)
	assert.NoError(t, err)

	assert.Equal(t, CONTINUE, c)
	assert.NotEqual(t, unpermitHost, req.URL.Host)
}

// test for AuthDirector
func TestAuthDirector(t *testing.T) {
	h := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "1", r.Header.Get("X-Ochanoco-UserID"))
		fmt.Fprintln(w, "Hello, client")
	}

	testDirector := func(proxy *OchanocoProxy, req *http.Request, context *gin.Context) (bool, error) {
		session := sessions.Default(context)
		session.Set("userId", "1")
		err := session.Save()
		assert.NoError(t, err)

		c, err := AuthDirector(proxy, req, context)

		assert.NoError(t, err)
		assert.Equal(t, CONTINUE, c)

		return CONTINUE, nil
	}

	DEFAULT_DIRECTORS = []OchanocoDirector{
		testDirector,
	}

	req, recorder, _, proxy := directorSample(t)
	mockServer, _ := setupMockServer(h, req, t)
	defer mockServer.Close()

	req.URL.Path = "/hello?hoge"

	proxy.Engine.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)
}

func TestAuthDirectorWithWhiteList(t *testing.T) {
	h := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}

	DEFAULT_DIRECTORS = []OchanocoDirector{
		AuthDirector,
	}

	req, recorder, _, proxy := directorSample(t)
	mockServer, _ := setupMockServer(h, req, t)
	defer mockServer.Close()

	proxy.Config.WhiteListPath = []string{
		"/hello",
	}

	proxy.Engine.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)
}

// test for AuthDirector
func TestAuthDirectorNoPermit(t *testing.T) {
	DEFAULT_DIRECTORS = []OchanocoDirector{
		AuthDirector,
	}

	req, recorder, _, proxy := directorSample(t)
	req.URL.Path = "/hello"

	proxy.Engine.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusUnauthorized, recorder.Result().StatusCode)
}

type TestResponseRecorder struct {
	*httptest.ResponseRecorder
	closeChannel chan bool
}

func (r *TestResponseRecorder) CloseNotify() <-chan bool {
	return r.closeChannel
}

func (r *TestResponseRecorder) closeClient() {
	r.closeChannel <- true
}

func CreateTestResponseRecorder() *TestResponseRecorder {
	return &TestResponseRecorder{
		httptest.NewRecorder(),
		make(chan bool, 1),
	}
}

func TestLogDirector(t *testing.T) {
	req, recorder, context, proxy := directorSample(t)

	before, err := proxy.Database.Client.RequestLog.Query().Count(proxy.Database.Ctx)
	assert.NoError(t, err)

	req.URL.Path = "/"

	LogDirector(proxy, req, context)

	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)

	after, err := proxy.Database.Client.RequestLog.Query().Count(proxy.Database.Ctx)
	assert.NoError(t, err)

	assert.Equal(t, before+1, after)

	all, err := proxy.Database.Client.RequestLog.Query().All(proxy.Database.Ctx)
	assert.NoError(t, err)

	requestLog := all[after-1]
	t.Log("--- HEADER ---")
	t.Log(requestLog.Headers)
}
