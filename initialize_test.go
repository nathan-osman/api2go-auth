package auth

import (
	"net/http"
	"testing"

	"github.com/manyminds/api2go"
)

type TestItem struct{}

func (t *TestItem) GetName() string {
	return "items"
}

func (t *TestItem) GetID() string {
	return "id"
}

type TestResource struct{}

func (t *TestResource) FindAll(req api2go.Request) (api2go.Responder, error) {
	return &api2go.Response{
		Res:  nil,
		Code: http.StatusOK,
	}, nil
}

var TestInitializeData = []struct {
	ShouldInitialize bool
	Status           int
}{
	{
		Status: http.StatusInternalServerError,
	},
	{
		ShouldInitialize: true,
		Status:           http.StatusOK,
	},
}

func TestInitialize(t *testing.T) {
	for _, v := range TestInitializeData {
		var (
			api = api2go.NewAPI("")
			h   = New(api, &TestAuthenticator{
				ShouldSucceed:      true,
				ShouldAuthenticate: true,
				ShouldInitialize:   v.ShouldInitialize,
			}, nil)
		)
		api.AddResource(&TestItem{}, &TestResource{})
		r, err := sendRequest(
			h,
			http.MethodPost,
			"/login",
			nil,
			nil,
			http.StatusOK,
		)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := sendRequest(
			h,
			http.MethodGet,
			"/items",
			r.Cookies(),
			nil,
			v.Status,
		); err != nil {
			t.Fatal(err)
		}
	}
}
