package db

import (
	"fmt"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	MissingDatabaseID int = 0
)

type Database struct {
	Lock       *sync.Mutex
	Connection *gorm.DB
}

func NewDatabase(connection *Connection) (*Database, error) {
	if err := connection.Parse(); err != nil {
		return nil, fmt.Errorf("unable to parse database connection; %w", err)
	}

	db, err := gorm.Open(postgres.Open(connection.String))
	if err != nil {
		return nil, err
	}

	return &Database{
		Lock:       &sync.Mutex{},
		Connection: db,
	}, nil
}

// Read is a generic function to read a model
// when no special logic is required.
func (database *Database) Read(id int, model interface{}) error {
	result := database.Connection.Find(model, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("unable to find id [%d]", id)
	}

	return nil
}

// FindBy is a generic function to read a model by
// which matches a field with a particular value.
func (database *Database) FindBy(field string, value interface{}, model interface{}) (*gorm.DB, error) {
	result := database.Connection.Where(map[string]interface{}{field: value}).First(model)
	if result.Error != nil {
		if result.RowsAffected == 0 {
			return result, nil
		}

		return result, result.Error
	}

	return result, nil
}

// Delete is a generic function to delete a model
// when no special logic is required.
func (database *Database) Delete(id int, model interface{}) error {
	result := database.Connection.Delete(model, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
