package auth

import (
	"net/http"
	"testing"

	"github.com/manyminds/api2go"
)

var TestLoginData = []struct {
	ShouldSucceed      bool
	ShouldAuthenticate bool
	Status             int
}{
	{
		Status: http.StatusInternalServerError,
	},
	{
		ShouldSucceed: true,
		Status:        http.StatusForbidden,
	},
	{
		ShouldSucceed:      true,
		ShouldAuthenticate: true,
		Status:             http.StatusOK,
	},
}

func TestLogin(t *testing.T) {
	for _, v := range TestLoginData {
		var (
			api = api2go.NewAPI("api")
			h   = New(api, &TestAuthenticator{
				ShouldSucceed:      v.ShouldSucceed,
				ShouldAuthenticate: v.ShouldAuthenticate,
			}, nil)
		)
		if _, err := sendRequest(
			h,
			http.MethodPost,
			"/login",
			nil,
			nil,
			v.Status,
		); err != nil {
			t.Fatal(err)
		}
	}
}
