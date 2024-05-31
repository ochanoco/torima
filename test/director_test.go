package test

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
	"github.com/ochanoco/torima/core"
	"github.com/ochanoco/torima/extension/directors"
	"github.com/ochanoco/torima/proxy"
	"github.com/stretchr/testify/assert"
)

func directorSample(t *testing.T) (*core.TorimaPackageContext[*http.Request], *TestResponseRecorder) {
	core.DB_TYPE = "sqlite3"
	core.DB_CONFIG = "../data/test.db?_fk=1"
	core.SECRET = "test_secret"

	recorder := CreateTestResponseRecorder()
	ginContext, r := gin.CreateTestContext(recorder)

	store := cookie.NewStore([]byte("test"))
	r.Use(sessions.Sessions("torima-session", store))

	db, err := core.InitDB(core.DB_CONFIG)
	assert.NoError(t, err)

	config, file, err := readTestConfig(t)
	assert.NoError(t, err)
	defer os.Remove(file.Name())

	proxy := core.NewOchancoProxy(r, proxy.DEFAULT_DIRECTORS, proxy.DEFAULT_MODIFY_RESPONSES, proxy.DEFAULT_PROXYWEB_PAGES, config, db)
	req := httptest.NewRequest("GET", "http://localhost:8080/", nil)

	ctx := core.TorimaPackageContext[*http.Request]{
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
	c, err := directors.BasicRoute("example.com", ctx)

	assert.NoError(t, err)
	assert.Equal(t, core.Keep, c)
	assert.Equal(t, "example.com", ctx.Target.URL.Host)
	assert.Equal(t, "http", ctx.Target.URL.Scheme)
	assert.Equal(t, core.SECRET, ctx.Target.Header.Get("X-Torima-Proxy-Token"))
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

	c, err := directors.ThirdPartyDirector(ctx)
	assert.NoError(t, err)
	assert.Equal(t, core.Keep, c)

	c, err = directors.ThirdPartyDirector(ctx)
	assert.NoError(t, err)

	assert.Equal(t, core.Keep, c)
	assert.Equal(t, host, ctx.Target.URL.Host)
}

// test for DefaultRouteDirector
func TestThirdPartyDirectorNoParmit(t *testing.T) {
	unpermitHost := "not-in-list.example.com"

	ctx, _ := directorSample(t)

	ctx.Target.URL.Path = "/torima/redirect/" + unpermitHost + "/"

	c, err := directors.ThirdPartyDirector(ctx)
	assert.NoError(t, err)
	assert.Equal(t, core.Keep, c)

	c, err = directors.ThirdPartyDirector(ctx)
	assert.NoError(t, err)

	assert.Equal(t, core.Keep, c)
	assert.NotEqual(t, unpermitHost, ctx.Target.URL.Host)
}

// test for AuthDirector
func TestAuthDirector(t *testing.T) {
	h := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "1", r.Header.Get("X-Torima-UserID"))
		fmt.Fprintln(w, "Hello, client")
	}

	testDirector := func(c *core.TorimaPackageContext[*http.Request]) (core.TorimaPackageStatus, error) {
		session := sessions.Default(c.GinContext)

		user := ninsho.LINE_USER{
			Sub: "1",
		}
		json, _ := json.Marshal(user)

		session.Set("user", string(json))
		err := session.Save()
		assert.NoError(t, err)

		status, err := directors.AuthDirector(c)

		assert.NoError(t, err)
		assert.Equal(t, core.Authed, status)

		return core.Authed, nil
	}

	proxy.DEFAULT_DIRECTORS = core.TorimaDirectors{
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

	proxy.DEFAULT_DIRECTORS = core.TorimaDirectors{
		directors.AuthDirector,
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
	proxy.DEFAULT_DIRECTORS = core.TorimaDirectors{
		directors.AuthDirector,
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

	directors.BeforeLogDirector(ctx)

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
