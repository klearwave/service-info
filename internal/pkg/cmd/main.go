package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/klearwave/service-info/internal/pkg/db"
	"github.com/klearwave/service-info/internal/pkg/server"
)

const (
	httpPort  = 8888
	httpsPort = 8443

	tlsCertPath = "/tls.crt"
	tlsKeyPath  = "/tls.key"
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

	var hasCert, hasKey bool
	if _, err := os.Stat(tlsCertPath); err == nil {
		hasCert = true
	}

	if _, err := os.Stat(tlsKeyPath); err == nil {
		hasKey = true
	}

	// use certificates if they exist, otherwise start the server without TLS
	if hasCert && hasKey {
		log.Printf("Starting HTTPS server on port %d", httpsPort)
		err = http.ListenAndServeTLS(
			fmt.Sprintf("0.0.0.0:%d", httpsPort),
			tlsCertPath,
			tlsKeyPath,
			server.Router,
		)
	} else {
		log.Printf("Starting HTTP server on port %d", httpPort)
		err = http.ListenAndServe(
			fmt.Sprintf("0.0.0.0:%d", httpPort),
			server.Router,
		)
	}

	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
