package auth

import (
	"errors"
	"net/http"

	"github.com/manyminds/api2go/jsonapi"
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
	b, err := jsonapi.Marshal(i)
	if err != nil {
		writeError(w, err, http.StatusInternalServerError)
		return
	}
	session, _ := a.store.Get(r, sessionName)
	session.Values[sessionData] = id
	session.Save(r, w)
	writeJSON(w, b, http.StatusOK)
}
