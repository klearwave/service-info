package v0

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"gorm.io/gorm"

	"github.com/klearwave/service-info/pkg/api"
	apierrors "github.com/klearwave/service-info/pkg/api/errors"
	"github.com/klearwave/service-info/pkg/api/model"
	"github.com/klearwave/service-info/pkg/db"
	"github.com/klearwave/service-info/pkg/utils/pointers"
)

const (
	VersionRegex = `^v?(\d+)\.(\d+)\.(\d+)(?:-([0-9A-Za-z.-]+))?(?:\+([0-9A-Za-z.-]+))?$`
)

// VersionBase is the base set of fields for all Version objects.
type VersionBase struct {
	Id *string `json:"id,omitempty" gorm:"primarykey" example:"v0.1.2" doc:"Version in semantic versioning format."`
}

// Version represents the full database schema for a Version.  The full schema is also used in responses.
type Version struct {
	model.Model
	VersionBase

	Stable *bool `json:"stable,omitempty" example:"false" doc:"Whether this is a stable version."`

	XVersion     *int    `json:"x_version,omitempty" example:"0" doc:"The X version of this release (e.g. 0; v0.1.2-prerelease.1 == x.y.z-build.metadata)."`
	YVersion     *int    `json:"y_version,omitempty" example:"1" doc:"The Y version of this release (e.g. 1; v0.1.2-prerelease.1 == x.y.z-build.metadata)."`
	ZVersion     *int    `json:"z_version,omitempty" example:"2" doc:"The Z version of this release (e.g. 2; v0.1.2-prerelease.1 == x.y.z-build.metadata)."`
	BuildVersion *string `json:"build_version,omitempty" example:"prerelease.1" doc:"The build version and metadata of this release (e.g. prerelease.1; v0.1.2-prerelease.1 == x.y.z-build.metadata)."`

	ContainerImages []*ContainerImage `json:"container_images,omitempty" gorm:"many2many:version_container_images;" doc:"Container images associated with this version."`
}

// BeforeCreate defines the before create logic for a specific version.  The BeforeCreate
// is used in conjunction with GORM as a trigger function that is called before
// inserting a record into the database.
func (version *Version) BeforeCreate(tx *gorm.DB) error {
	// set the x, y and z versions for the version
	if err := version.Parse(); err != nil {
		return fmt.Errorf("%s; %w", apierrors.ErrInvalidVersion, err)
	}

	return nil
}

// Create handles the create request for a version model.
func (version *Version) Create(database *db.Database) *api.Result {
	return api.Create(database, version)
}

// Read handles the read request for a version model.
func (version *Version) Read(database *db.Database) *api.Result {
	apiResult := &api.Result{}

	defer func() {
		apiResult.Object = version
	}()

	// query the database and return any errors
	databaseResult := database.Connection.Where(map[string]interface{}{"id": *version.Id}).
		Preload("ContainerImages").
		First(version)

	if databaseResult.Error != nil {
		if errors.Is(databaseResult.Error, gorm.ErrRecordNotFound) {
			apiResult.NotFoundError(nil, *version.Id, version)

			return apiResult
		}

		apiResult.UnknownError(databaseResult.Error)

		return apiResult
	}

	if version.Id == nil || *version.Id == "" {
		apiResult.NotFoundError(nil, *version.Id, version)

		return apiResult
	}

	return apiResult
}

// Versions represents a list of versions.
type Versions []Version

// List handles the list request for a set of versions.
func (versions *Versions) List(database *db.Database) *api.Result {
	apiResult := &api.Result{}

	defer func() {
		apiResult.Object = versions
	}()

	databaseResult := database.Connection.Preload("ContainerImages").Find(versions)
	if databaseResult.Error != nil {
		apiResult.UnknownError(databaseResult.Error)

		return apiResult
	}

	return apiResult
}

// Delete handles the delete request for a version model.
func (version *Version) Delete(database *db.Database) *api.Result {
	return api.Delete(database, *version.Id, version)
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
	matches := regexp.MustCompile(VersionRegex).FindStringSubmatch(versionId)
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
