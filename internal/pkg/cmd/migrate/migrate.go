package migrate

import (
	"context"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"

	"github.com/klearwave/service-info/internal/pkg/db"
)

const migrateExample = `
service migrate
`

type input struct {
	Directory  string
	Connection *db.Config
}

func NewCommand() *cobra.Command {
	commandInput := &input{
		Connection: &db.Config{},
	}

	// create the command
	command := &cobra.Command{
		Use:     "migrate",
		Short:   "Run database migrations",
		Long:    `Run database migrations`,
		RunE:    func(_ *cobra.Command, args []string) error { return migrate(args, commandInput) },
		Example: migrateExample,
	}

	command.Flags().StringVarP(&commandInput.Directory, "directory", "d", "./migrations", "directory with migration files")
	command.Flags().StringVarP(&commandInput.Connection.DatabaseName, "db-name", "n", "postgres", "Database name to perform migrations against")
	command.Flags().StringVar(&commandInput.Connection.Host, "db-host", "localhost", "Database host where migration database resides")
	command.Flags().IntVar(&commandInput.Connection.Port, "db-port", 5432, "Port which database is running on")
	command.Flags().StringVarP(&commandInput.Connection.Username, "db-username", "u", "postgres", "Username which has access to the database")
	command.Flags().StringVarP(&commandInput.Connection.Password, "db-password", "p", "postgres", "Password of user which was access to the database")

	return command
}

//nolint:forbidigo
func migrate(args []string, commandInput *input) error {
	if len(args) < 1 {
		return fmt.Errorf("goose: missing command")
	}

	command, connection := args[0], commandInput.Connection

	if err := connection.Parse(); err != nil {
		return err
	}

	db, err := goose.OpenDBWithDriver("postgres", connection.String)
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v", err)
		}
	}()

	arguments := []string{}
	if len(args) > 1 {
		arguments = append(arguments, args[1:]...)
	}

	ctx := context.Background()
	if err := goose.RunContext(ctx, command, db, commandInput.Directory, arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}

	return nil
}
