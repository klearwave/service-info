package main

// const (
// 	httpPort  = 8888
// 	httpsPort = 8443

// 	tlsCertPath = "/tls.crt"
// 	tlsKeyPath  = "/tls.key"
// )

// func main() {
// 	// create the server
// 	server, err := server.NewServer(&db.Config{})
// 	if err != nil {
// 		panic(err)
// 	}

// 	// ensure we have database connectivity
// 	if err := server.Database.Wait(30); err != nil {
// 		panic(err)
// 	}

// 	// register our routes once we confirm database connectivity is established
// 	server.RegisterRoutes()

// 	var hasCert, hasKey bool
// 	if _, err := os.Stat(tlsCertPath); err == nil {
// 		hasCert = true
// 	}

// 	if _, err := os.Stat(tlsKeyPath); err == nil {
// 		hasKey = true
// 	}

// 	// use certificates if they exist, otherwise start the server without TLS
// 	if hasCert && hasKey {
// 		log.Printf("Starting HTTPS server on port %d", httpsPort)
// 		err = http.ListenAndServeTLS(
// 			fmt.Sprintf("0.0.0.0:%d", httpsPort),
// 			tlsCertPath,
// 			tlsKeyPath,
// 			server.Router,
// 		)
// 	} else {
// 		log.Printf("Starting HTTP server on port %d", httpPort)
// 		err = http.ListenAndServe(
// 			fmt.Sprintf("0.0.0.0:%d", httpPort),
// 			server.Router,
// 		)
// 	}

// 	if err != nil {
// 		log.Fatalf("Failed to start server: %v", err)
// 	}
// }

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/klearwave/service-info/internal/pkg/cmd/generate"
	"github.com/klearwave/service-info/internal/pkg/cmd/health"
	"github.com/klearwave/service-info/internal/pkg/cmd/migrate"
	"github.com/klearwave/service-info/internal/pkg/cmd/run"
	"github.com/klearwave/service-info/internal/pkg/cmd/version"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(command *cobra.Command) {
	err := command.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// main executes the main program loop.
func main() {
	root := &cobra.Command{
		Use:   "service",
		Short: "Manage service",
		Long:  `Manage service`,
	}

	// add version subcommand for printing version information
	root.AddCommand(version.NewCommand())

	// add other subcommands
	root.AddCommand(generate.NewCommand())
	root.AddCommand(migrate.NewCommand())
	root.AddCommand(run.NewCommand())
	root.AddCommand(health.NewCommand())

	Execute(root)
}
