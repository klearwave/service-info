package models

import (
	"fmt"
	"regexp"
	"strconv"

	"gorm.io/gorm"

	"github.com/klearwave/service-info/pkg/errors"
	"github.com/klearwave/service-info/pkg/models"
	"github.com/klearwave/service-info/pkg/utils/pointers"
)

const (
	versionRegex = `^v?(\d+)\.(\d+)\.(\d+)(?:-([0-9A-Za-z.-]+))?(?:\+([0-9A-Za-z.-]+))?$`
)

// VersionBase is the base set of fields for all Version objects.
type VersionBase struct {
	Id *string `json:"id,omitempty" gorm:"primarykey" example:"v0.1.2" doc:"Version in semantic versioning format."`
}

// VersionRequestBody is the base set of inputs required to create a Version.
type VersionRequestBody struct {
	VersionBase

	ContainerImages []*ContainerImageRequestBody `json:"container_images,omitempty" gorm:"many2many:version_container_images;" doc:"Container images associated with this version."`
}

// TableName defines the table name for the request.
func (request VersionRequestBody) TableName() string {
	return "versions"
}

// ToResponse converts a request object to a response object.
func (request *VersionRequestBody) ToResponse() *Version {
	response := &Version{}

	response.VersionBase = request.VersionBase

	if len(request.ContainerImages) < 1 {
		return response
	}

	response.ContainerImages = make([]*ContainerImage, len(request.ContainerImages))

	for i := range request.ContainerImages {
		response.ContainerImages[i] = request.ContainerImages[i].ToResponse()
	}

	return response
}

// Version represents the full database schema for a Version.  The full schema is also used in responses.
type Version struct {
	models.Model
	VersionBase

	Stable *bool `json:"stable,omitempty" example:"false" doc:"Whether this is a stable version."`

	XVersion     *int    `json:"x_version,omitempty" example:"0" doc:"The X version of this release (e.g. 0; v0.1.2-prerelease.1 == x.y.z-build.metadata)."`
	YVersion     *int    `json:"y_version,omitempty" example:"1" doc:"The Y version of this release (e.g. 1; v0.1.2-prerelease.1 == x.y.z-build.metadata)."`
	ZVersion     *int    `json:"z_version,omitempty" example:"2" doc:"The Z version of this release (e.g. 2; v0.1.2-prerelease.1 == x.y.z-build.metadata)."`
	BuildVersion *string `json:"build_version,omitempty" example:"prerelease.1" doc:"The build version and metadata of this release (e.g. prerelease.1; v0.1.2-prerelease.1 == x.y.z-build.metadata)."`

	ContainerImages []*ContainerImage `json:"container_images,omitempty" gorm:"many2many:version_container_images;" doc:"Container images associated with this version."`
}

// VersionRequestCreate represents the request when creating a version.
type VersionRequestCreate struct {
	models.Authorization

	Body VersionRequestBody
}

// VersionRequestGet represents the request when getting a version.
type VersionRequestGet struct {
	Id string `path:"id"`
}

// VersionRequestDelete represents the request when deleting a version.
type VersionRequestDelete struct {
	models.Authorization

	Id string `path:"id"`
}

// VersionRequestList represents the request when listing all versions.
type VersionRequestList struct{}

// VersionResponse is the response returned that includes the Version struct as the response body.
type VersionResponse struct {
	Body   Version
	Status int
}

// VersionsResponse is the response returned when multiple versions are requested.
type VersionsResponse struct {
	Body   []Version
	Status int
}

// BeforeCreate defines the before create logic for a specific version.  The BeforeCreate
// is used in conjunction with GORM as a trigger function that is called before
// inserting a record into the database.
func (version *Version) BeforeCreate(tx *gorm.DB) error {
	if err := version.validate(); err != nil {
		return fmt.Errorf("%s; %w", errors.ErrInvalidVersion, err)
	}

	// set the x, y and z versions for the version
	if err := version.Parse(); err != nil {
		return fmt.Errorf("%s; %w", errors.ErrInvalidVersion, err)
	}

	return nil
}

// Parse sets the major, minor, bugfix and build versions for a specific version.  It also
// performs some basic mutations and validations againt the set version.
func (version *Version) Parse() error {
	versionId := *version.Id

	// add 'v' prefix if missing
	if versionId[0] != 'v' {
		version.Id = pointers.FromString("v" + versionId)
	}

	// regular expression for semantic versioning with optional 'v' prefix
	matches := regexp.MustCompile(versionRegex).FindStringSubmatch(versionId)
	if matches == nil {
		return fmt.Errorf("version not in semantic versioning format: %s", versionId)
	}

	// extract major, minor, patch and build versions
	subVersions := make([]int, 3)

	for i := range []int{0, 1, 2} {
		number, err := strconv.Atoi(matches[i+1])
		if err != nil {
			return err
		}

		subVersions[i] = number
	}

	// finally store the versions
	version.XVersion = pointers.Int(subVersions[0])
	version.YVersion = pointers.Int(subVersions[1])
	version.ZVersion = pointers.Int(subVersions[2])

	// set the optional build version
	if matches[4] != "" {
		version.BuildVersion = pointers.FromString(matches[4])
	} else {
		// if this is not a build version, e.g. v1.2.3-alpha.1, then
		// we can consider it to be a stable version
		version.Stable = pointers.Bool(true)
	}

	return nil
}

// validate validates that a specific version is valid.
func (version *Version) validate() error {
	if version == nil {
		return errors.ErrMissingVersionObject
	}

	if version.Id == nil {
		return errors.ErrMissingVersionParameterVersionId
	}

	re, err := regexp.Compile(versionRegex)
	if err != nil {
		return fmt.Errorf("invalid regex pattern [%s]: %w", versionRegex, err)
	}

	if !re.MatchString(*version.Id) {
		return fmt.Errorf("input [%s] does not match the required pattern [%s]",
			*version.Id,
			versionRegex,
		)
	}

	return nil
}
