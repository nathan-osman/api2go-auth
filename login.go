package auth

import (
	"net/http"
)

func (a *Auth) login(w http.ResponseWriter, r *http.Request) {
	id, i, err := a.authenticator.Authenticate(r)
	if err != nil {
		writeError(w, err, http.StatusForbidden)
		return
	}
	session, _ := a.store.Get(r, sessionName)
	session.Values[sessionData] = id
	session.Save(r, w)
	writeJSON(w, i, http.StatusOK)
}
