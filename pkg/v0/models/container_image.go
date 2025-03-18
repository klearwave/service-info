package models

import (
	"errors"
	"fmt"
	"regexp"

	"gorm.io/gorm"

	apierrors "github.com/klearwave/service-info/pkg/errors"
	"github.com/klearwave/service-info/pkg/models"
	"github.com/klearwave/service-info/pkg/utils/pointers"
)

const (
	containerImageRegex = `^((?P<registry>[^/]+?)/)?(?P<image>[^:@]+(?:/[^:@]+)*)(?::(?P<tag>[^@]+))?(?:@sha256:(?P<sha>[a-zA-Z0-9]+))?$`
)

// ContainerImageBase is the base set of fields for all ContainerImage objects.
type ContainerImageBase struct {
	Image      *string `json:"image,omitempty" example:"ghcr.io/klearwave/service-info:latest" doc:"Full container image including the registry, repository and tag."`
	SHA256Sum  *string `json:"sha256sum,omitempty" example:"2d4b92db6941294f731cfe7aeca336eb8dba279171c0e6ceda32b9f018f8429d" doc:"SHA256 sum of the container image."`
	CommitHash *string `json:"commit_hash,omitempty" example:"631af50a8bbc4b5e69dab77d51a3a1733550fe8d" doc:"Commit hash related to the image."`
}

// ContainerImageRequestBody is the base set of inputs required to create a ContainerImage.
type ContainerImageRequestBody struct {
	ContainerImageBase
}

// TableName defines the table name for the request.
func (request ContainerImageRequestBody) TableName() string {
	return "container_images"
}

// FromRequest converts a request object to a response object.
func (request *ContainerImageRequestBody) ToResponse() *ContainerImage {
	response := &ContainerImage{}
	response.ContainerImageBase = request.ContainerImageBase

	return response
}

// ContainerImage represents the full database schema for a ContainerImage.  The full schema is also used in responses.
type ContainerImage struct {
	models.Model

	ContainerImageBase

	ImageRegistry *string `json:"image_registry,omitempty" gorm:"default:docker.io" example:"ghcr.io" doc:"Container image registry without the image name, tag or sha256sum information."`
	ImageName     *string `json:"image_name,omitempty" example:"klearwave/service-info" doc:"Container image name without the registry, tag or sha256sum information."`
	ImageTag      *string `json:"image_tag,omitempty" gorm:"default:latest" example:"v0.1.2" doc:"Container image tag without the registry, image name or sha256 information."`
}

// EqualTo compares another container image for equality and returns
// a bool if they are not equal.
func (containerImage *ContainerImage) EqualTo(compared *ContainerImage) (bool, error) {
	for field, values := range map[string][]*string{
		"sha256sum":      {containerImage.SHA256Sum, compared.SHA256Sum},
		"commit_hash":    {containerImage.CommitHash, compared.CommitHash},
		"image":          {containerImage.Image, compared.Image},
		"image_registry": {containerImage.ImageRegistry, compared.ImageRegistry},
		"image_name":     {containerImage.ImageName, compared.ImageName},
		"image_tag":      {containerImage.ImageTag, compared.ImageTag},
	} {
		if !pointers.EqualString(values[0], values[1]) {
			return false, fmt.Errorf(
				"field [%s] not equal requested [%s] != existing [%s]; %w",
				field,
				pointers.ToString(values[0]),
				pointers.ToString(values[1]),
				apierrors.ErrMismatchParameter,
			)
		}
	}

	return true, nil
}

// BeforeCreate defines the before create logic for a specific container image.  The BeforeCreate
// is used in conjunction with GORM as a trigger function that is called before
// inserting a record into the database.
func (containerImage *ContainerImage) BeforeCreate(tx *gorm.DB) error {
	if err := containerImage.Parse(); err != nil {
		return fmt.Errorf("%s; %w", apierrors.ErrInvalidContainerImage, err)
	}

	existing := &ContainerImage{}

	err := tx.Where("image = ?", containerImage.Image).First(existing).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}

	if err != nil {
		return fmt.Errorf("unable to lookup existing container image by name; %w", err)
	}

	containerImage.Id = existing.Id

	equal, err := containerImage.EqualTo(existing)
	if err != nil {
		return err
	}

	if !equal {
		return apierrors.ErrMismatchParameter
	}

	return nil
}

// Parse sets specific container image metadata such as the image path and the
// image tag.
func (containerImage *ContainerImage) Parse() error {
	if containerImage == nil {
		return apierrors.ErrMissingContainerImageObject
	}

	if containerImage.Image == nil {
		return apierrors.ErrInvalidContainerImage
	}

	image := *containerImage.Image

	// regular expression for splitting container images
	re, err := regexp.Compile(containerImageRegex)
	if err != nil {
		return fmt.Errorf("invalid container image pattern [%s]: %w", containerImageRegex, err)
	}

	matches := re.FindStringSubmatch(image)
	if matches == nil || len(matches) != 6 {
		return fmt.Errorf("invalid container image [%s] for pattern [%s]", image, containerImageRegex)
	}

	// validate that we do not have a mismatch between the full image sha256 (if provided) and the
	// value from the input and set the appropriate value.
	if containerImage.SHA256Sum == nil {
		containerImage.SHA256Sum = pointers.FromString(matches[5])
	} else {
		if *containerImage.SHA256Sum == "" {
			containerImage.SHA256Sum = pointers.FromString(matches[5])
		} else {
			if matches[5] != "" && matches[5] != *containerImage.SHA256Sum {
				return fmt.Errorf("found sha256sum value [%s] but parsed [%s]; %w",
					*containerImage.SHA256Sum,
					matches[5],
					apierrors.ErrMismatchParameter,
				)
			}
		}
	}

	// set the other parsed values.  note that no comparison between what is set and
	// what is parsed needs to happen as we do not provide these values via the request.
	containerImage.ImageRegistry = pointers.FromString(matches[2])
	containerImage.ImageName = pointers.FromString(matches[3])
	containerImage.ImageTag = pointers.FromString(matches[4])

	return containerImage.validate()
}

// validate validates that a specific container image is valid.
func (containerImage *ContainerImage) validate() error {
	missingErrors := []error{}

	for err, value := range map[error]*string{
		apierrors.ErrMissingContainerImageParameterImage:      containerImage.Image,
		apierrors.ErrMissingContainerImageParameterSHA256Sum:  containerImage.SHA256Sum,
		apierrors.ErrMissingContainerImageParameterCommitHash: containerImage.CommitHash,
	} {
		if value == nil {
			missingErrors = append(missingErrors, err)
		}
	}

	if len(missingErrors) > 0 {
		return fmt.Errorf("%v", missingErrors)
	}

	return nil
}

// ContainerImageRequestCreate represents the request when creating a version.
type ContainerImageRequestCreate struct {
	Body ContainerImageRequestBody
}

// ContainerImageRequestGet represents the request when getting a version.
type ContainerImageRequestGet struct {
	Id int `path:"id"`
}

// ContainerImageRequestDelete represents the request when deleting a version.
type ContainerImageRequestDelete struct {
	Id int `path:"id"`
}

// ContainerImageRequestList represents the request when listing all versions.
type ContainerImageRequestList struct{}

// ContainerImageResponse is the response returned that includes the ContainerImage struct as the response body.
type ContainerImageResponse struct {
	Body   ContainerImage
	Status int
}

// ContainerImagesResponse is the response returned when multiple versions are requested.
type ContainerImagesResponse struct {
	Body   []ContainerImage
	Status int
}
