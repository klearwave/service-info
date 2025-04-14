package run

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"

	"github.com/klearwave/service-info/internal/pkg/db"
	"github.com/klearwave/service-info/internal/pkg/server"
)

const (
	httpPort  = 8888
	httpsPort = 8443

	tlsCertPath = "/tls.crt"
	tlsKeyPath  = "/tls.key"
)

const runExample = `
service run
`

func NewCommand() *cobra.Command {
	// create the command
	command := &cobra.Command{
		Use:     "run",
		Short:   "Run the service",
		Long:    `Run the service`,
		RunE:    run,
		Example: runExample,
	}

	return command
}

//nolint:forbidigo
func run(cmd *cobra.Command, args []string) error {
	// create the server
	server, err := server.NewServer()
	if err != nil {
		return fmt.Errorf("failed setting up server: %w", err)
	}
	server.RegisterRoutes()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// initialize the server
	if err := server.Init(&db.Config{}); err != nil {
		return err
	}

	// start the server and block waiting for cancel
	if err := server.Start(); err != nil {
		return err
	}

	<-ctx.Done()

	// gracefully stop the server
	shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	return server.Stop(shutdownCtx)
}
