package auth

import (
	"net/http"
	"testing"
)

var TestInitializeData = []struct {
	ShouldInitialize bool
	Cookies          bool
	Status           int
}{
	{
		Status: http.StatusForbidden,
	},
	{
		Cookies: true,
		Status:  http.StatusInternalServerError,
	},
	{
		ShouldInitialize: true,
		Cookies:          true,
		Status:           http.StatusOK,
	},
}

func TestInitialize(t *testing.T) {
	for _, v := range TestInitializeData {
		h := createAPI(true, true, v.ShouldInitialize)
		c, err := login(h, http.StatusOK)
		if err != nil {
			t.Fatal(err)
		}
		if !v.Cookies {
			c = nil
		}
		if _, err := sendRequest(
			h,
			http.MethodGet,
			"/items",
			c,
			nil,
			v.Status,
		); err != nil {
			t.Fatal(err)
		}
	}
}
