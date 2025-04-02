package service

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
)

// TODO: WIP generic handling of CRUD operations for a web service.

// Create handles the creation of a new resource in the database for a web service.
func (service *Service) Create(ctx context.Context, request Request) (*Response, error) {
	service.Database.Lock.Lock()
	defer service.Database.Lock.Unlock()

	// authorize the request
	authorized, err := request.IsAuthorized()
	if err != nil || !authorized {
		return nil, huma.Error401Unauthorized(
			"request is unauthorized",
			err,
		)
	}

	// validate the request
	valid, err := request.IsValid()
	if err != nil || !valid {
		return nil, huma.Error422UnprocessableEntity(
			"request is invalid",
			err,
		)
	}

	// send the request to the database
	response, err := request.ToModel().Create(service.Database.Connection)
	if err != nil {
		return nil, err
	}

	return response, nil
}
