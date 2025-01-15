package routes

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	v0 "github.com/klearwave/service-info/pkg/v0"
	"github.com/klearwave/service-info/pkg/v0/models"
)

const (
	DefaultVersionsPath = "versions"
	DefaultVersionsTag  = "Versions"
)

// CreateVersion defines the routing and subsequent specification for POST to /api/v0/versions.
func CreateVersion(input *models.Version) huma.Operation {
	return huma.Operation{
		OperationID:   "createVersion",
		Summary:       "Create a new Version.",
		Description:   "Create a new Version.",
		Method:        http.MethodPost,
		DefaultStatus: http.StatusCreated,
		Path:          v0.PathFor(DefaultVersionsPath),
		Tags:          []string{DefaultVersionsTag},
	}
}

// GetVersion defines the routing and subsequent specification for GET to /api/v0/versions/{version_id}.
func GetVersion(input *models.Version) huma.Operation {
	return huma.Operation{
		OperationID:   "getVersion",
		Summary:       "Get specific version information.",
		Description:   "Get specific version information.",
		Method:        http.MethodGet,
		DefaultStatus: http.StatusOK,
		Path:          v0.PathFor(DefaultVersionsPath) + "/{version_id}",
		Parameters: []*huma.Param{
			{
				Name:        "version_id",
				Description: "Version ID to get.",
				Example:     "v0.1.0",
				In:          "path",
				Required:    true,
			},
		},
		Tags: []string{DefaultVersionsTag},
	}
}

// GetVersions defines the routing and subsequent specification for GET to /api/v0/versions.
func GetVersions(input *models.Version) huma.Operation {
	return huma.Operation{
		OperationID:   "getVersions",
		Summary:       "Get all version information.",
		Description:   "Get all version information.",
		Method:        http.MethodGet,
		DefaultStatus: http.StatusOK,
		Path:          v0.PathFor(DefaultVersionsPath),
		Tags:          []string{DefaultVersionsTag},
	}
}

// DeleteVersion defines the routing and subsequent specification for DELETE to /api/v0/versions/{version_id}.
func DeleteVersion(input *models.Version) huma.Operation {
	return huma.Operation{
		OperationID:   "deleteVersion",
		Summary:       "Delete a specific version.",
		Description:   "Delete a specific version.",
		Method:        http.MethodDelete,
		DefaultStatus: http.StatusOK,
		Path:          v0.PathFor(DefaultVersionsPath) + "/{version_id}",
		Parameters: []*huma.Param{
			{
				Name:        "version_id",
				Description: "Version ID to delete.",
				Example:     "v0.1.0",
				In:          "path",
				Required:    true,
			},
		},
		Tags: []string{DefaultVersionsTag},
	}
}
