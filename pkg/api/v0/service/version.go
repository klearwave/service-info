package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"gorm.io/gorm"

	modelsv0 "github.com/klearwave/service-info/pkg/api/v0/models"
	apierrors "github.com/klearwave/service-info/pkg/errors"
	"github.com/klearwave/service-info/pkg/models"
)

// CreateVersion defines the service for creating a new version and storing in a database.
func (service *Service) CreateVersion(ctx context.Context, request *modelsv0.VersionRequestCreate) (*modelsv0.VersionResponse, error) {
	service.Database.Lock.Lock()
	defer service.Database.Lock.Unlock()

	// create the version
	response := request.Body.ToResponse()

	result := service.Database.Connection.Create(response)
	if result.Error != nil {
		return nil, result.Error
	}

	return &modelsv0.VersionResponse{
		Body: *response,
	}, nil
}

// GetVersion defines the service for retrieving a specific version from the database.
func (service *Service) GetVersion(ctx context.Context, request *modelsv0.VersionRequestGet) (*modelsv0.VersionResponse, error) {
	service.Database.Lock.Lock()
	defer service.Database.Lock.Unlock()

	if request.Id == "" {
		return nil, apierrors.ErrMissingVersionParameterVersionId
	}

	response := &modelsv0.VersionResponse{Body: modelsv0.Version{}}

	result := service.Database.Connection.Where(map[string]interface{}{"id": request.Id}).
		Preload("ContainerImages").
		First(&response.Body)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			response.Status = http.StatusNotFound

			return response, huma.Error404NotFound(
				fmt.Sprintf("unable to find version with id: [%s]", request.Id),
				result.Error,
			)
		}

		response.Status = http.StatusInternalServerError

		return response, result.Error
	}

	if *response.Body.Id == "" {
		response.Status = http.StatusNotFound

		return response, huma.Error404NotFound(
			fmt.Sprintf("found version with missing id: [%s]", request.Id),
			result.Error,
		)
	}

	return response, nil
}

// GetVersions defines the service for retrieving all versions from a database.
func (service *Service) GetVersions(ctx context.Context, request *modelsv0.VersionRequestList) (*modelsv0.VersionsResponse, error) {
	service.Database.Lock.Lock()
	defer service.Database.Lock.Unlock()

	versions := []modelsv0.Version{}

	result := service.Database.Connection.Preload("ContainerImages").Find(&versions)
	if result.Error != nil {
		return nil, result.Error
	}

	return &modelsv0.VersionsResponse{
		Body: versions,
	}, nil
}

// DeleteVersion defines the service for deleting a version from a database.
func (service *Service) DeleteVersion(ctx context.Context, request *modelsv0.VersionRequestDelete) (*models.DeleteResponse, error) {
	service.Database.Lock.Lock()
	defer service.Database.Lock.Unlock()

	version := modelsv0.Version{}

	if request.Id == "" {
		return nil, apierrors.ErrMissingVersionParameterVersionId
	}

	if _, err := service.Database.FindBy("id", request.Id, &version); err != nil {
		return nil, err
	}

	result := service.Database.Connection.Delete(version)
	if result.Error != nil {
		return nil, result.Error
	}

	response := &models.DeleteResponse{}
	response.Body.Message = "Delete Success"

	return response, nil
}

// GetVersionContainerImages defines the service for getting all container images belonging to a specific version.
func (service *Service) GetVersionContainerImages(ctx context.Context, request *modelsv0.VersionRequestGet) (*modelsv0.ContainerImagesResponse, error) {
	service.Database.Lock.Lock()
	defer service.Database.Lock.Unlock()

	if request.Id == "" {
		return nil, apierrors.ErrMissingVersionParameterVersionId
	}

	version := &modelsv0.Version{
		VersionBase: modelsv0.VersionBase{
			Id: &request.Id,
		},
	}

	response := &modelsv0.ContainerImagesResponse{}

	err := service.Database.Connection.
		Model(version).
		Association("ContainerImages").
		Find(&response.Body)

	if err != nil {
		response.Status = http.StatusInternalServerError

		return nil, err
	}

	return response, nil
}
