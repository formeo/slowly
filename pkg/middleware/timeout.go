package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

func CheckTimeoutValue(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		done := make(chan bool, 1)
		ctx, cancelFunc := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancelFunc()
		go func() {
			next.ServeHTTP(w, r)
			close(done)
		}()

		select {
		case <-done:
			return
		case <-ctx.Done():
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "timeout too long"})
		}
	})

}
