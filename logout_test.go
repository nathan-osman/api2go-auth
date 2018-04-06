package auth

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"testing"
)

func TestLogout(t *testing.T) {
	u, err := url.Parse("http://example.com")
	if err != nil {
		t.Fatal(err)
	}
	j, err := cookiejar.New(nil)
	if err != nil {
		t.Fatal(err)
	}
	h := createAPI(true, true, true)
	c, err := login(h, http.StatusOK)
	if err != nil {
		t.Fatal(err)
	}
	j.SetCookies(u, c)
	if err := findAll(h, j.Cookies(u), http.StatusOK); err != nil {
		t.Fatal(err)
	}
	r, err := sendRequest(
		h,
		http.MethodPost,
		"/logout",
		c,
		http.StatusOK,
	)
	if err != nil {
		t.Fatal(err)
	}
	j.SetCookies(u, r.Cookies())
	if err := findAll(h, j.Cookies(u), http.StatusForbidden); err != nil {
		t.Fatal(err)
	}
}
