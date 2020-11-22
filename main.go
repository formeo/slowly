package main

import (
	"encoding/json"
	"io"
	"log"
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
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var postTimeout PostTimeout
		_ = json.NewDecoder(r.Body).Decode(&postTimeout)
		if postTimeout.Timeout > 5000 {

			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			io.WriteString(w, `{"error":"timeout too long"}`)
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
func startServer(){
	router := mux.NewRouter()
	router.HandleFunc("/heathcheck", heathcheck).Methods("GET")
	router.HandleFunc("/api/slow", returnPost).Methods("POST")
	router.Use(Middleware)
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}

func main() {
	startServer()

}
