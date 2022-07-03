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

func delCookie(name string, src []*http.Cookie) []*http.Cookie {
	var dest []*http.Cookie

	for _, v := range src {
		if v.Name != name {
			dest = append(dest, v)
		}
	}

	return dest
}
