package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
	"testing"
)

func TestRemoveToken(t *testing.T) {
	msg := "hello"
	token := "hello"
	rp := httputil.ReverseProxy{
		Director:       director,
		ModifyResponse: modifyResponse,
	}

	Directors = append(Directors, RemoveToken)

	t.Run("enc/dec token", func(t *testing.T) {
		simpleServ := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
			if req.URL.Query().Get("token") == token {
				t.Errorf("Error: removing token (1)")
			}

			_, err := req.Cookie("token")

			if err != nil {
				t.Errorf("Error: removing token (2)")
			}

			fmt.Fprintln(writer, msg)
		}))

		simpleDirector := func(req *http.Request) {
			url, _ := url.Parse(simpleServ.URL)
			req.URL.Scheme = url.Scheme
			req.URL.Host = url.Host
			req.URL.Path = "/"
		}

		Directors = append(Directors, simpleDirector)
		targetServ := httptest.NewServer(&rp)

		url := fmt.Sprintf("%s?token=%s", targetServ.URL, token)
		req, err := http.NewRequest(http.MethodGet, url, nil)

		if err != nil {
			t.Error(err)
		}

		req.AddCookie(&http.Cookie{Name: "token", Value: token})

		_, err = new(http.Client).Do(req)
		if err != nil {
			t.Error(err)
		}
	})
}
