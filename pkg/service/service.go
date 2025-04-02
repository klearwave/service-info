package service

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/klearwave/service-info/pkg/db"
	"gorm.io/gorm"
)

// Request is an interface that represents a generic request object to be used by a service.
type Request interface {
	IsValid() (bool, error)
	IsAuthorized() (bool, error)
	ToModel() Model
}

// Response is an object that represents a generic response object to be used by a service.
type Response struct {
	Status int
	Body   interface{}
	Error  huma.StatusError
}

// Model is an interface that represents a generic model object to be used by a service.
type Model interface {
	Create(db *gorm.DB) (*Response, error)
	Read(db *gorm.DB) (*Response, error)
	Update(db *gorm.DB) (*Response, error)
	Delete(db *gorm.DB) (*Response, error)
}

// Service is an object that represents a service that interacts with a database.
type Service struct {
	Database *db.Database
}

func NewService(db *db.Database) *Service {
	return &Service{
		Database: db,
	}
}
