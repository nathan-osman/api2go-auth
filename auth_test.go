package auth

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/manyminds/api2go"
)

type TestItem struct{}

func (t *TestItem) GetName() string { return "items" }
func (t *TestItem) GetID() string   { return "id" }

type TestResource struct{}

func (t *TestResource) FindAll(api2go.Request) (api2go.Responder, error) {
	return &api2go.Response{
		Res:  nil,
		Code: http.StatusOK,
	}, nil
}

func createAPI(shouldSucceed, shouldAuthenticate, shouldInitialize bool) *Auth {
	var (
		api = api2go.NewAPI("")
		h   = New(api, &TestAuthenticator{
			ShouldSucceed:      shouldSucceed,
			ShouldAuthenticate: shouldAuthenticate,
			ShouldInitialize:   shouldInitialize,
		}, nil)
	)
	api.AddResource(&TestItem{}, &TestResource{})
	return h
}

func sendRequest(h http.Handler, method, target string, cookies []*http.Cookie, code int) (*http.Response, error) {
	var (
		w = httptest.NewRecorder()
		r = httptest.NewRequest(method, target, nil)
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

func login(h http.Handler, code int) ([]*http.Cookie, error) {
	r, err := sendRequest(
		h,
		http.MethodPost,
		"/login",
		nil,
		code,
	)
	if err != nil {
		return nil, err
	}
	return r.Cookies(), nil
}

func findAll(h http.Handler, cookies []*http.Cookie, code int) error {
	_, err := sendRequest(
		h,
		http.MethodGet,
		"/items",
		cookies,
		code,
	)
	return err
}
