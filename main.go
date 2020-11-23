package main

import (
	router2 "github.com/formeo/slowly/pkg/router"
	"github.com/formeo/slowly/pkg/server"
)

func main() {
	router := router2.NewRouter()
	server.NewServer(router)
}
