package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/ochanoco/ninsho"
	"github.com/stretchr/testify/assert"
)

func directorSample(t *testing.T) (*TorimaPackageContext[*http.Request], *TestResponseRecorder) {
	DB_TYPE = "sqlite3"
	DB_CONFIG = "../data/test.db?_fk=1"
	SECRET = "test_secret"

	recorder := CreateTestResponseRecorder()
	ginContext, r := gin.CreateTestContext(recorder)

	store := cookie.NewStore([]byte("test"))
	r.Use(sessions.Sessions("torima-session", store))

	db, err := InitDB(DB_CONFIG)
	assert.NoError(t, err)

	config, file, err := readTestConfig(t)
	assert.NoError(t, err)
	defer os.Remove(file.Name())

	proxy := NewOchancoProxy(r, DEFAULT_DIRECTORS, DEFAULT_MODIFY_RESPONSES, DEFAULT_PROXYWEB_PAGES, config, db)
	req := httptest.NewRequest("GET", "http://localhost:8080/", nil)

	ctx := TorimaPackageContext[*http.Request]{
		GinContext: ginContext,
		Proxy:      &proxy,
		Target:     req,
	}

	return &ctx, recorder
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
	ctx, _ := directorSample(t)
	c, err := RouteDirector("example.com", ctx)

	assert.NoError(t, err)
	assert.Equal(t, Keep, c)
	assert.Equal(t, "example.com", ctx.Target.URL.Host)
	assert.Equal(t, "http", ctx.Target.URL.Scheme)
	assert.Equal(t, SECRET, ctx.Target.Header.Get("X-Torima-Proxy-Token"))
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

	ctx, _ := directorSample(t)

	ctx.Target.URL.Path = "/torima/redirect/" + host

	ctx.Proxy.Config.ProtectionScope = []string{host}

	c, err := ThirdPartyDirector(ctx)
	assert.NoError(t, err)
	assert.Equal(t, Keep, c)

	c, err = ThirdPartyDirector(ctx)
	assert.NoError(t, err)

	assert.Equal(t, Keep, c)
	assert.Equal(t, host, ctx.Target.URL.Host)
}

// test for DefaultRouteDirector
func TestThirdPartyDirectorNoParmit(t *testing.T) {
	unpermitHost := "not-in-list.example.com"

	ctx, _ := directorSample(t)

	ctx.Target.URL.Path = "/torima/redirect/" + unpermitHost + "/"

	c, err := ThirdPartyDirector(ctx)
	assert.NoError(t, err)
	assert.Equal(t, Keep, c)

	c, err = ThirdPartyDirector(ctx)
	assert.NoError(t, err)

	assert.Equal(t, Keep, c)
	assert.NotEqual(t, unpermitHost, ctx.Target.URL.Host)
}

// test for AuthDirector
func TestAuthDirector(t *testing.T) {
	h := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "1", r.Header.Get("X-Torima-UserID"))
		fmt.Fprintln(w, "Hello, client")
	}

	testDirector := func(c *TorimaPackageContext[*http.Request]) (TorimaPackageStatus, error) {
		session := sessions.Default(c.GinContext)

		user := ninsho.LINE_USER{
			Sub: "1",
		}
		json, _ := json.Marshal(user)

		session.Set("user", string(json))
		err := session.Save()
		assert.NoError(t, err)

		status, err := AuthDirector(c)

		assert.NoError(t, err)
		assert.Equal(t, Authed, status)

		return Authed, nil
	}

	DEFAULT_DIRECTORS = TorimaDirectors{
		testDirector,
	}

	ctx, recorder := directorSample(t)
	mockServer, _ := setupMockServer(h, ctx.Target, t)
	defer mockServer.Close()

	ctx.Target.URL.Path = "/hello?hoge"

	ctx.Proxy.Engine.ServeHTTP(recorder, ctx.Target)
	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)
}

func TestAuthDirectorWithWhiteList(t *testing.T) {
	h := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}

	DEFAULT_DIRECTORS = TorimaDirectors{
		AuthDirector,
	}

	ctx, recorder := directorSample(t)
	mockServer, _ := setupMockServer(h, ctx.Target, t)
	defer mockServer.Close()

	ctx.Proxy.Config.SkipAuthList = []string{
		"/hello",
	}

	ctx.Proxy.Engine.ServeHTTP(recorder, ctx.Target)
	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)
}

// test for AuthDirector
func TestAuthDirectorNoPermit(t *testing.T) {
	DEFAULT_DIRECTORS = TorimaDirectors{
		AuthDirector,
	}

	ctx, recorder := directorSample(t)
	ctx.Target.URL.Path = "/hello"

	ctx.Proxy.Engine.ServeHTTP(recorder, ctx.Target)
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
	ctx, recorder := directorSample(t)

	before, err := ctx.Proxy.Database.Client.RequestLog.Query().Count(ctx.Proxy.Database.Ctx)
	assert.NoError(t, err)

	ctx.Target.URL.Path = "/"

	BeforeLogDirector(ctx)

	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)

	after, err := ctx.Proxy.Database.Client.RequestLog.Query().Count(ctx.Proxy.Database.Ctx)
	assert.NoError(t, err)

	assert.Equal(t, before+1, after)

	all, err := ctx.Proxy.Database.Client.RequestLog.Query().All(ctx.Proxy.Database.Ctx)
	assert.NoError(t, err)

	requestLog := all[after-1]
	t.Log("--- HEADER ---")
	t.Log(requestLog.Headers)

	assert.Equal(t, "before", requestLog.Flag)
}
