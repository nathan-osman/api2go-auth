package auth

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
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
// authenticator for requests. The secret key allows for persistent sessions
// and may be set to nil to disable the feature.
func New(api *api2go.API, authenticator Authenticator, secretKey []byte) *Auth {
	if secretKey == nil {
		secretKey = securecookie.GenerateRandomKey(32)
	}
	var (
		r = mux.NewRouter()
		a = &Auth{
			api:           api,
			authenticator: authenticator,
			router:        r,
			store:         sessions.NewCookieStore(secretKey),
		}
	)
	r.HandleFunc("/login", a.login).Methods(http.MethodPost)
	r.HandleFunc("/logout", a.logout).Methods(http.MethodPost)
	r.HandleFunc("/", a.initialize)
	return a
}

// ServeHTTP responds to requests for API resources.
func (a *Auth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
