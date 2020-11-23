package middleware

import (
	"encoding/json"
	"github.com/formeo/slowly/pkg/models"
	"io"
	"net/http"
)

func CheckTimeoutValue(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var postTimeout models.PostTimeout
		_ = json.NewDecoder(r.Body).Decode(&postTimeout)
		if postTimeout.Timeout > 5000 {

			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			_, err := io.WriteString(w, `{"error":"timeout too long"}`)
			if err != nil {
				panic(err)
			}
			return

		}

		h.ServeHTTP(w, r)
	})
}
