package core

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// test for RouteDirector
func TestRouteDirector(t *testing.T) {
	DB_TYPE = "sqlite3"
	DB_CONFIG = "../data/test.db?_fk=1"
	SECRET = "test_secret"

	proxy, err := ProxyServer()
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)

	req := httptest.NewRequest("GET", "http://localhost:8080/", nil)

	c, err := RouteDirector("example.com", proxy, req, context)

	assert.NoError(t, err)
	assert.Equal(t, CONTINUE, c)
	assert.Equal(t, "example.com", req.URL.Host)
	assert.Equal(t, "https", req.URL.Scheme)
	assert.Equal(t, "ochanoco", req.Header.Get("User-Agent"))
	assert.Equal(t, SECRET, req.Header.Get("X-Ochanoco-Proxy-Token"))
}
