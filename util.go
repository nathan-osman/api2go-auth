package auth

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/manyminds/api2go"
)

func writeJSON(w http.ResponseWriter, b []byte, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Content-Length", strconv.Itoa(len(b)))
	w.WriteHeader(status)
	w.Write(b)
}

func writeError(w http.ResponseWriter, err error, status int) {
	b, err := json.Marshal(
		api2go.NewHTTPError(
			err,
			err.Error(),
			status,
		),
	)
	if err != nil {
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}
	writeJSON(w, b, status)
}
