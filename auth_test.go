package auth

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

func sendRequest(h http.Handler, method, target string, cookies []*http.Cookie, body io.Reader, code int) (*http.Response, error) {
	var (
		w = httptest.NewRecorder()
		r = httptest.NewRequest(method, target, body)
	)
	for _, c := range cookies {
		r.AddCookie(c)
	}
	h.ServeHTTP(w, r)
	if w.Code != code {
		return nil, fmt.Errorf("%d != %d", w.Code, code)
	}
	return w.Result(), nil
}
