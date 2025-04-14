package db

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var (
	ErrMissingDatabaseHost     = errors.New("missing database host configuration")
	ErrMissingDatabaseName     = errors.New("missing database name configuration")
	ErrMissingDatabaseUsername = errors.New("missing database username configuration")
	ErrMissingDatabasePassword = errors.New("missing database password configuration")
)

const (
	defaultDatabasePort = 5432

	envDatabaseHost     = "DB_HOST"
	envDatabasePort     = "DB_PORT"
	envDatabaseName     = "DB_NAME"
	envDatabaseUsername = "DB_USERNAME"
	envDatabasePassword = "DB_PASSWORD"
)

// Config represents an object which contains the database connection parameters.
type Config struct {
	Host         string
	Port         int
	DatabaseName string
	Username     string
	Password     string

	String string
}

// Parse parses a connection object into a database string.  It uses environment variables
// if it cannot find parameters on the object.
func (connection *Config) Parse() error {
	if connection.Port == 0 {
		if os.Getenv(envDatabasePort) != "" {
			port, err := strconv.Atoi(os.Getenv(envDatabasePort))
			if err != nil {
				return fmt.Errorf("unable to convert %s to integer", os.Getenv(envDatabasePort))
			}

			connection.Port = port
		} else {
			connection.Port = defaultDatabasePort
		}
	}

	if connection.Host == "" {
		if os.Getenv(envDatabaseHost) != "" {
			connection.Host = os.Getenv(envDatabaseHost)
		} else {
			return ErrMissingDatabaseHost
		}
	}

	if connection.DatabaseName == "" {
		if os.Getenv(envDatabaseName) != "" {
			connection.DatabaseName = os.Getenv(envDatabaseName)
		} else {
			return ErrMissingDatabaseName
		}
	}

	if connection.Username == "" {
		if os.Getenv(envDatabaseUsername) != "" {
			connection.Username = os.Getenv(envDatabaseUsername)
		} else {
			return ErrMissingDatabaseUsername
		}
	}

	if connection.Password == "" {
		if os.Getenv(envDatabasePassword) != "" {
			connection.Password = os.Getenv(envDatabasePassword)
		} else {
			return ErrMissingDatabasePassword
		}
	}

	connection.String = fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		connection.Host,
		connection.Port,
		connection.Username,
		connection.Password,
		connection.DatabaseName,
	)

	return nil
}
