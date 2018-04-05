package auth

import (
	"errors"
	"net/http"
)

var errAuthRequired = errors.New("authentication required")

func (a *Auth) initialize(w http.ResponseWriter, r *http.Request) {
	session, _ := a.store.Get(r, sessionName)
	i, _ := session.Values[sessionData]
	if i == nil {
		writeError(w, errAuthRequired, http.StatusForbidden)
		return
	}
	r, err := a.authenticator.Initialize(r, i)
	if err != nil {
		writeError(w, err, http.StatusInternalServerError)
		return
	}
	a.api.Handler().ServeHTTP(w, r)
}
