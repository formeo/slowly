package main

import (
	"encoding/json"
	"net/http"
	"time"
)
import "github.com/gorilla/mux"

type PostTimeout struct {
	Timeout int `json:"timeout"`
}

type Response struct {
	Status string `json:"status"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func heathcheck(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"heathcheck": "ok"}`))
}

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var postTimeout PostTimeout
		_ = json.NewDecoder(r.Body).Decode(&postTimeout)
		if postTimeout.Timeout > 5000 {

			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{"error": "timeout too long"})
			return

		}

		h.ServeHTTP(w, r)
	})
}

func returnPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var postTimeout PostTimeout
	_ = json.NewDecoder(r.Body).Decode(&postTimeout)
	time.Sleep(time.Duration(postTimeout.Timeout) * time.Millisecond)
	response := Response{
		Status: "ok",
	}

	jsonResponse, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(jsonResponse)

	if err != nil {
		panic(err)
	}

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/heathcheck", heathcheck).Methods("GET")
	router.HandleFunc("/api/slow", returnPost).Methods("POST")
	router.Use(Middleware)
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		panic(err)
	}

}
