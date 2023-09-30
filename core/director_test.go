package core

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// test for RouteDirector
func TestRouteDirector(t *testing.T) {
	DB_TYPE = "sqlite3"
	DB_CONFIG = "../data/test.db?_fk=1"
	SECRET = "test_secret"

	proxy, err := ProxyServer()

	if err != nil {
		t.Fatalf("ProxyServer() is failed: %v", err)
	}

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)

	req := httptest.NewRequest("GET", "http://localhost:8080/", nil)

	c, err := RouteDirector("example.com", proxy, req, context)

	if err != nil {
		t.Fatalf("RouteDirector() is failed: %v", err)
	}

	if c != CONTINUE {
		t.Fatalf("RouteDirector() is failed: no-continued (%v)", c)
	}

	if req.URL.Host != "example.com" {
		t.Fatalf("RouteDirector() is failed: wrong host (%v)", req.Host)
	}

	if req.URL.Scheme != "https" {
		t.Fatalf("RouteDirector() is failed: wrong scheme (%v)", req.URL.Scheme)
	}

	agent := req.Header.Get("User-Agent")
	if agent != "ochanoco" {
		t.Fatalf("RouteDirector() is failed: %v", err)
	}

	token := req.Header.Get("X-Ochanoco-Proxy-Token")
	if token != SECRET {
		t.Fatalf("RouteDirector() is failed: %v", token)
	}
}
