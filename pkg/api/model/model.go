package model

import (
	"time"

	apierrors "github.com/klearwave/service-info/pkg/api/errors"
)

// Model represents the base model that is used for all other models.
type Model struct {
	CreatedAt *time.Time `json:"created_at,omitempty" example:"YYYY-MM-DDTHH:MM:SSZ" doc:"Object creation time in RFC 3339 format."`
	UpdatedAt *time.Time `json:"updated_at,omitempty" example:"YYYY-MM-DDTHH:MM:SSZ" doc:"Object last updated time in RFC 3339 format."`
}

// ModelWithId represents the base model that is used for all other models when an ID is used as
// a primary key.  If this model is not used, it is assumed that the underlying model defines
// its own primary key.
type ModelWithId struct {
	Model

	Id int `json:"id,omitempty" gorm:"primarykey" example:"1" doc:"Database ID of the stored object."`
}

// StringFetcher represents a model that uses a string with an 'id' parameter to fetch a model
// from the database.  It is used to inject into other models so that validation may be reused
// across multiple models.
type StringFetcher struct {
	Id string `path:"id"`
}

// IntegerFetcher represents a model that uses an integer with an 'id' parameter to fetch a model
// from the database.  It is used to inject into other models so that validation may be reused
// across multiple models.
type IntegerFetcher struct {
	Id int `path:"id"`
}

// StringFetcher.IsValid validates that a string fetcher is valid.
func (fetcher StringFetcher) IsValid() (bool, error) {
	if fetcher.Id == "" {
		return false, apierrors.ErrMissingParameterId
	}

	return true, nil
}

// IntegerFetcher.IsValid validates that a integer fetcher is valid.  In this case, we will never
// have a database ID of less than 1 (the first entry always has a positive number).
func (fetcher IntegerFetcher) IsValid() (bool, error) {
	if fetcher.Id < 1 {
		return false, apierrors.ErrMissingParameterId
	}

	return true, nil
}
