package models

import "time"

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

// DeleteResponse represents the generic delete response message for all models.
type DeleteResponse struct {
	Body struct {
		Message string
	}
}
