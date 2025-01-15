package db

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"github.com/klearwave/service-info/pkg/v0/models"
)

var (
	ErrMissingVersionId = errors.New("missing version_id parameter")
	ErrInvalidVersionId = errors.New("invalid version id")
)

// CreateVersion defines the service for creating a new version and storing in a database.
func (database *Database) CreateVersion(ctx context.Context, request *models.VersionRequestCreate) (*models.VersionResponse, error) {
	database.Lock.Lock()
	defer database.Lock.Unlock()

	version := &models.Version{}
	version.VersionBase = request.Body

	if version.VersionId == "" {
		return nil, ErrMissingVersionId
	}

	// validate that we have a valid semantic version
	if err := validateSemanticVersioning(version.VersionId); err != nil {
		return nil, fmt.Errorf("%s; %w", ErrInvalidVersionId, err)
	}

	// add 'v' prefix if missing
	if version.VersionId[0] != 'v' {
		version.VersionId = "v" + version.VersionId
	}

	// create the version
	if err := database.Create(version); err != nil {
		return nil, err
	}

	return &models.VersionResponse{
		Body: *version,
	}, nil
}

// GetVersion defines the service for retrieving a specific version from the database.
func (database *Database) GetVersion(ctx context.Context, request *models.VersionRequestGet) (*models.VersionResponse, error) {
	database.Lock.Lock()
	defer database.Lock.Unlock()

	response := &models.VersionResponse{Body: models.Version{}}

	if request.VersionID == "" {
		return nil, ErrMissingVersionId
	}

	if _, err := database.FindBy("version_id", request.VersionID, &response.Body); err != nil {
		return nil, err
	}

	if response.Body.VersionId == "" {
		response.Status = http.StatusNotFound

		return response, nil
	}

	return response, nil
}

// GetVersions defines the service for retrieving all versions from a database.
func (database *Database) GetVersions(ctx context.Context, request *models.VersionRequestList) (*models.VersionsResponse, error) {
	database.Lock.Lock()
	defer database.Lock.Unlock()

	versions := []models.Version{}

	result := database.Connection.Find(&versions)
	if result.Error != nil {
		return nil, result.Error
	}

	return &models.VersionsResponse{
		Body: versions,
	}, nil
}

// Delete defines the service for deleting a version from a database.
func (database *Database) DeleteVersion(ctx context.Context, request *models.VersionRequestDelete) (*models.DeleteResponse, error) {
	database.Lock.Lock()
	defer database.Lock.Unlock()

	version := models.Version{}

	if request.VersionID == "" {
		return nil, ErrMissingVersionId
	}

	if _, err := database.FindBy("version_id", request.VersionID, &version); err != nil {
		return nil, err
	}

	if err := database.Delete(version.Id, version); err != nil {
		return nil, err
	}

	response := &models.DeleteResponse{}
	response.Body.Message = "Delete Success"

	return response, nil
}

// validateSemanticVersioning checks if the given string matches semantic versioning syntax.
// If the 'v' prefix is missing, it adds the prefix.
func validateSemanticVersioning(version string) error {
	// regular expression for semantic versioning with optional 'v' prefix
	semverRegex := `^v?(\d+)\.(\d+)\.(\d+)(?:-([\da-zA-Z-]+(?:\.[\da-zA-Z-]+)*))?(?:\+([\da-zA-Z-]+(?:\.[\da-zA-Z-]+)*))?$`

	re := regexp.MustCompile(semverRegex)
	if !re.MatchString(version) {
		return fmt.Errorf("invalid semantic version: %s", version)
	}

	return nil
}
