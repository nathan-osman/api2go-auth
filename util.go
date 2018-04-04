package auth

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// writeJSON outputs a JSON response for the provided data.
func writeJSON(w http.ResponseWriter, i interface{}, status int) {
	b, err := json.Marshal(i)
	if err != nil {
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Content-Length", strconv.Itoa(len(b)))
	w.WriteHeader(status)
	w.Write(b)
}

// writeError outputs a JSON response for an error.
func writeError(w http.ResponseWriter, err error, status int) {
	writeJSON(w, map[string]string{"error": err.Error()}, status)
}
