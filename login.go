package auth

import (
	"errors"
	"net/http"
)

var errAuthFailed = errors.New("authentication failed")

func (a *Auth) login(w http.ResponseWriter, r *http.Request) {
	id, i, err := a.authenticator.Authenticate(r)
	if err != nil {
		writeError(w, err, http.StatusInternalServerError)
		return
	}
	if id == nil {
		writeError(w, errAuthFailed, http.StatusForbidden)
		return
	}
	session, _ := a.store.Get(r, sessionName)
	session.Values[sessionData] = id
	session.Save(r, w)
	writeJSON(w, i, http.StatusOK)
}
