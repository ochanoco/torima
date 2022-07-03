package main

import (
	"errors"
	"net/http"
)

func parseCookies(rawCookies string) []*http.Cookie {
	header := http.Header{}
	header.Add("Cookie", rawCookies)
	request := http.Request{Header: header}
	cookies := request.Cookies()

	return cookies
}

func getCookie(name string, cookies []*http.Cookie) (*http.Cookie, error) {
	err := errors.New("proxy/utils: cookie not found")

	for _, v := range cookies {
		if v.Name == name {
			return v, nil
		}
	}

	return nil, err
}
