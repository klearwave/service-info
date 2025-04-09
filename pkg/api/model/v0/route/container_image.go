package route

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	v0 "github.com/klearwave/service-info/pkg/api/model/v0"
)

const (
	DefaultContainerImagesPath     = "container_images"
	DefaultContainerImagesGroupTag = "ContainerImages"
)

// CreateContainerImage defines the routing and subsequent specification for POST to /api/v0/container_images.
func CreateContainerImage() huma.Operation {
	return huma.Operation{
		OperationID:   "createContainerImage",
		Summary:       "Create a new container image.",
		Description:   "Create a new container image.",
		Method:        http.MethodPost,
		DefaultStatus: http.StatusCreated,
		Path:          v0.PathFor(DefaultContainerImagesPath),
		Tags:          containerImagesGroupTags(),
	}
}

// GetContainerImage defines the routing and subsequent specification for GET to /api/v0/container_images/{id}.
func GetContainerImage() huma.Operation {
	return huma.Operation{
		OperationID:   "getContainerImage",
		Summary:       "Get specific container image information.",
		Description:   "Get specific container image information.",
		Method:        http.MethodGet,
		DefaultStatus: http.StatusOK,
		Path:          v0.PathFor(DefaultContainerImagesPath) + "/{id}",
		Parameters: []*huma.Param{
			{
				Name:        "id",
				Description: "Database ID of container image to get.",
				Example:     1,
				In:          "path",
				Required:    true,
			},
		},
		Tags: containerImagesGroupTags(),
	}
}

// ListContainerImages defines the routing and subsequent specification for GET to /api/v0/container_images.
func ListContainerImages() huma.Operation {
	return huma.Operation{
		OperationID:   "listContainerImages",
		Summary:       "List all container image information.",
		Description:   "List all container image information.",
		Method:        http.MethodGet,
		DefaultStatus: http.StatusOK,
		Path:          v0.PathFor(DefaultContainerImagesPath),
		Tags:          containerImagesGroupTags(),
	}
}

// ListContainerImageVersions defines the routing and subsequent specification for GET to /api/v0/container_images/{id}/versions.
func ListContainerImageVersions() huma.Operation {
	return huma.Operation{
		OperationID:   "listContainerImageVersions",
		Summary:       "List versions for a specific container image.",
		Description:   "List versions for a specific container image.",
		Method:        http.MethodGet,
		DefaultStatus: http.StatusOK,
		Path:          v0.PathFor(DefaultContainerImagesPath) + "/{id}/versions",
		Parameters: []*huma.Param{
			{
				Name:        "id",
				Description: "Container Image ID to get versions for.",
				Example:     1,
				In:          "path",
				Required:    true,
			},
		},
		Tags: containerImagesGroupTags(),
	}
}

// DeleteContainerImage defines the routing and subsequent specification for DELETE to /api/v0/container_images/{container_image_id}.
func DeleteContainerImage() huma.Operation {
	return huma.Operation{
		OperationID:   "deleteContainerImage",
		Summary:       "Delete a specific container image.",
		Description:   "Delete a specific container image.",
		Method:        http.MethodDelete,
		DefaultStatus: http.StatusOK,
		Path:          v0.PathFor(DefaultContainerImagesPath) + "/{id}",
		Parameters: []*huma.Param{
			{
				Name:        "id",
				Description: "Database ID of container image to delete.",
				Example:     1,
				In:          "path",
				Required:    true,
			},
		},
		Tags: containerImagesGroupTags(),
	}
}

// containerImagesGroupTags defines the group tags that group similar APIs in the documentation site.
func containerImagesGroupTags() []string {
	return []string{
		DefaultContainerImagesGroupTag,
		v0.DefaultGroupTag,
	}
}
