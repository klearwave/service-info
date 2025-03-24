package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	"github.com/klearwave/service-info/pkg/db"
	apierrors "github.com/klearwave/service-info/pkg/errors"
	"github.com/klearwave/service-info/pkg/models"
	modelsv0 "github.com/klearwave/service-info/pkg/v0/models"
)

// CreateContainerImage defines the service for creating a new container image and storing in a database.
func (service *Service) CreateContainerImage(ctx context.Context, request *modelsv0.ContainerImageRequestCreate) (*modelsv0.ContainerImageResponse, error) {
	service.Database.Lock.Lock()
	defer service.Database.Lock.Unlock()

	containerImage := request.Body.ToResponse()

	// create the container image
	result := service.Database.Connection.Create(containerImage)
	if result.Error != nil {
		return nil, result.Error
	}

	return &modelsv0.ContainerImageResponse{
		Body: *containerImage,
	}, nil
}

// GetContainerImage defines the service for retrieving a specific container image from the database.
func (service *Service) GetContainerImage(ctx context.Context, request *modelsv0.ContainerImageRequestGet) (*modelsv0.ContainerImageResponse, error) {
	service.Database.Lock.Lock()
	defer service.Database.Lock.Unlock()

	response := &modelsv0.ContainerImageResponse{Body: modelsv0.ContainerImage{}}

	if _, err := service.Database.FindBy("id", request.Id, &response.Body); err != nil {
		return nil, err
	}

	if response.Body.Id == db.MissingDatabaseID {
		response.Status = http.StatusNotFound

		return response, huma.Error404NotFound(fmt.Sprintf("unable to find container image with id: [%d]", response.Body.Id))
	}

	return response, nil
}

// GetContainerImages defines the service for retrieving all container images from a database.
func (service *Service) GetContainerImages(ctx context.Context, request *modelsv0.ContainerImageRequestList) (*modelsv0.ContainerImagesResponse, error) {
	service.Database.Lock.Lock()
	defer service.Database.Lock.Unlock()

	containerImages := []modelsv0.ContainerImage{}

	result := service.Database.Connection.Find(&containerImages)
	if result.Error != nil {
		return nil, result.Error
	}

	return &modelsv0.ContainerImagesResponse{
		Body: containerImages,
	}, nil
}

// DeleteContainerImage defines the service for deleting a container image from a database.
func (service *Service) DeleteContainerImage(ctx context.Context, request *modelsv0.ContainerImageRequestDelete) (*models.DeleteResponse, error) {
	service.Database.Lock.Lock()
	defer service.Database.Lock.Unlock()

	containerImage := modelsv0.ContainerImage{}

	if request.Id == db.MissingDatabaseID {
		return nil, apierrors.ErrMissingVersionParameterId
	}

	if _, err := service.Database.FindBy("id", request.Id, &containerImage); err != nil {
		return nil, err
	}

	if err := service.Database.Delete(containerImage.Id, containerImage); err != nil {
		return nil, err
	}

	response := &models.DeleteResponse{}
	response.Body.Message = "Delete Success"

	return response, nil
}

// GetContainerImageVersions defines the service for deleting a container image from a database.
func (service *Service) GetContainerImageVersions(ctx context.Context, request *modelsv0.ContainerImageRequestGet) (*modelsv0.VersionsResponse, error) {
	service.Database.Lock.Lock()
	defer service.Database.Lock.Unlock()

	if request.Id == 0 {
		return nil, apierrors.ErrMissingVersionParameterVersionId
	}

	containerImage := &modelsv0.ContainerImage{
		ModelWithId: models.ModelWithId{
			Id: request.Id,
		},
	}

	response := &modelsv0.VersionsResponse{}

	err := service.Database.Connection.
		Model(containerImage).
		Association("Versions").
		Find(&response.Body)

	if err != nil {
		response.Status = http.StatusInternalServerError

		return nil, err
	}

	return response, nil
}
