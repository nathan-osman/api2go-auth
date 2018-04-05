package auth

import "net/http"

// Authenticator provides methods for authenticating requests.
type Authenticator interface {

	// Authenticate determines if the request represents a valid login attempt.
	// The first return value is a unique identifier stored with the session
	// and used for initializing requests. The second return value is encoded
	// as JSON and returned to the client. The third return value is used if an
	// error occurs during authentication (a bad password, for example).
	Authenticate(r *http.Request) (interface{}, interface{}, error)

	// Initialize prepares an authenticated request for processing. Typically,
	// this involves setting a value on the request's context based on the
	// provided session object (returned by Authenticate).
	Initialize(r *http.Request, i interface{})
}
