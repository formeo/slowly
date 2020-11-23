package middleware

import (
	"context"
	"io"
	"net/http"
	"time"
)

func CheckTimeoutValue(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		done := make(chan bool)
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
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			_, err := io.WriteString(w, `{"error":"timeout too long"}`)
			if err != nil {
				panic(err)
			}
		}
	})

}
