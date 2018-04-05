package auth

import (
	"errors"
	"net/http"
)

const TestID = "test"

type TestType struct{}

func (t *TestType) GetID() string {
	return TestID
}

var errTest = errors.New("test")

type TestAuthenticator struct {
	ShouldSucceed      bool
	ShouldAuthenticate bool
	ShouldInitialize   bool
}

func (t *TestAuthenticator) Authenticate(r *http.Request) (interface{}, interface{}, error) {
	if !t.ShouldSucceed {
		return nil, nil, errTest
	}
	if !t.ShouldAuthenticate {
		return nil, nil, nil
	}
	return TestID, &TestType{}, nil
}

func (t *TestAuthenticator) Initialize(r *http.Request, i interface{}) (*http.Request, error) {
	if !t.ShouldInitialize {
		return nil, errTest
	}
	return r, nil
}
