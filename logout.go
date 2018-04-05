package auth

import (
	"net/http"
)

func (a *Auth) logout(w http.ResponseWriter, r *http.Request) {
	session, _ := a.store.Get(r, sessionName)
	session.Values[sessionData] = nil
	session.Save(r, w)
}
