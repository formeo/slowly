package router

import (
	"github.com/formeo/slowly/pkg/middleware"
	"github.com/formeo/slowly/pkg/views"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/healthCheck", views.HealthCheck).Methods("GET")
	router.HandleFunc("/api/slow", views.ReturnPostWithParamTimeout).Methods("POST")
	router.Use(middleware.CheckTimeoutValue)
	return router
}
