package auth

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manyminds/api2go"
)

// Authenticator provides methods for authenticating requests.
type Authenticator interface {

	// Authenticate determines if the request represents a valid login attempt
	// and returns an object to associate with the session or an error if the
	// login attempt is invalid.
	Authenticate(r *http.Request) (interface{}, error)

	// Initialize prepares an authenticated request for processing. Typically,
	// this involves setting a value on the request's context based on the
	// provided session object (set by Authenticate).
	Initialize(r *http.Request, i interface{})
}

// Auth provides a handler to authenticate requests. In addition to serving the
// routes provided by the API, two additional routes are added for logging in
// and logging out.
type Auth struct {
	api    *api2go.API
	a11r   Authenticator
	router *mux.Router
}

// New creates a new handler for the provided API using the provided
// authenticator for requests.
func New(api *api2go.API, a11r Authenticator) *Auth {
	var (
		r = mux.NewRouter()
		a = &Auth{
			api:    api,
			a11r:   a11r,
			router: r,
		}
	)
	r.HandleFunc("/login", a.login)
	r.HandleFunc("/logout", a.logout)
	r.Handle("/", api.Handler())
	return a
}

// ServeHTTP responds to requests for API resources.
func (a *Auth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO
}
