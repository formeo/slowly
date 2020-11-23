package server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func NewServer(router *mux.Router) {
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
