package main

import (
	"net/http"

	"github.com/klearwave/service-info/pkg/db"
	"github.com/klearwave/service-info/pkg/server"
)

func main() {
	// create the server
	server, err := server.NewServer(&db.Connection{})
	if err != nil {
		panic(err)
	}

	// ensure we have database connectivity
	if err := server.Database.Wait(30); err != nil {
		panic(err)
	}

	// register our routes once we confirm database connectivity is established
	server.RegisterRoutes()

	// start the server
	http.ListenAndServe("0.0.0.0:8888", server.Router)
}
