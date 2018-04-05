package auth

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/manyminds/api2go"
)

const (
	sessionName = "name"
	sessionData = "data"
)

// Auth provides a handler to authenticate requests. In addition to serving the
// routes provided by the API, two additional routes are added for logging in
// and logging out.
type Auth struct {
	api           *api2go.API
	authenticator Authenticator
	router        *mux.Router
	store         *sessions.CookieStore
}

// New creates a new handler for the provided API using the provided
// authenticator for requests.
func New(api *api2go.API, authenticator Authenticator, secretKey string) *Auth {
	var (
		r = mux.NewRouter()
		a = &Auth{
			api:           api,
			authenticator: authenticator,
			router:        r,
			store:         sessions.NewCookieStore([]byte(secretKey)),
		}
	)
	r.HandleFunc("/login", a.login).Methods(http.MethodPost)
	r.HandleFunc("/logout", a.logout).Methods(http.MethodPost)
	r.Handle("/", api.Handler())
	return a
}

// ServeHTTP responds to requests for API resources.
func (a *Auth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
