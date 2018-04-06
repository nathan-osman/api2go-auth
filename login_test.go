package auth

import (
	"net/http"
	"testing"
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
		h := createAPI(v.ShouldSucceed, v.ShouldAuthenticate, false)
		_, err := login(h, v.Status)
		if err != nil {
			t.Fatal(err)
		}
	}
}
