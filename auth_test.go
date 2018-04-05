package auth

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

func sendRequest(h http.Handler, method, target string, body io.Reader, code int) (*http.Response, error) {
	var (
		w = httptest.NewRecorder()
		r = httptest.NewRequest(method, target, body)
	)
	h.ServeHTTP(w, r)
	if w.Code != code {
		return nil, fmt.Errorf("%d != %d", w.Code, code)
	}
	return w.Result(), nil
}
