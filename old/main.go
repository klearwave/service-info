package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	apiv0 "github.com/klearwave/service-info/pkg/api/v0"
	serverv0 "github.com/klearwave/service-info/pkg/server/v0"
)

func main() {
	// create a type that satisfies the `api.ServerInterface`, which contains an implementation of every operation from the generated code
	server := serverv0.NewServer()

	r := gin.Default()

	apiv0.RegisterHandlers(r, server)

	// And we serve HTTP until the world ends.

	s := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8080",
	}

	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}
