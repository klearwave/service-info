package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"gorm.io/gorm"

	apierrors "github.com/klearwave/service-info/pkg/errors"
	"github.com/klearwave/service-info/pkg/models"
	modelsv0 "github.com/klearwave/service-info/pkg/v0/models"
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

	if request.VersionID == "" {
		return nil, apierrors.ErrMissingVersionParameterId
	}

	response := &modelsv0.VersionResponse{Body: modelsv0.Version{}}

	result := service.Database.Connection.Where(map[string]interface{}{"version_id": request.VersionID}).
		Preload("ContainerImages").
		First(&response.Body)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			response.Status = http.StatusNotFound

			return response, huma.Error404NotFound(
				fmt.Sprintf("unable to find version with version_id: [%s]", request.VersionID),
				result.Error,
			)
		}

		response.Status = http.StatusInternalServerError

		return response, result.Error
	}

	if *response.Body.VersionId == "" {
		response.Status = http.StatusNotFound

		return response, huma.Error404NotFound(
			fmt.Sprintf("found version with missing version_id: [%s]", request.VersionID),
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

	if request.VersionID == "" {
		return nil, apierrors.ErrMissingVersionParameterId
	}

	if _, err := service.Database.FindBy("version_id", request.VersionID, &version); err != nil {
		return nil, err
	}

	if err := service.Database.Delete(version.Id, version); err != nil {
		return nil, err
	}

	response := &models.DeleteResponse{}
	response.Body.Message = "Delete Success"

	return response, nil
}

// GetVersionContainerImages defines the service for getting all container images belonging to a specific version.
func (service *Service) GetVersionContainerImages(ctx context.Context, request *modelsv0.VersionRequestGet) (*modelsv0.ContainerImagesResponse, error) {
	service.Database.Lock.Lock()
	defer service.Database.Lock.Unlock()

	var containerImages []modelsv0.ContainerImage

	response := &modelsv0.ContainerImagesResponse{}

	err := service.Database.Connection.Model(&modelsv0.ContainerImage{}).Preload("Versions").Find(&containerImages).Error
	if err != nil {
		response.Status = http.StatusInternalServerError

		return response, nil
	}

	return response, nil
}
