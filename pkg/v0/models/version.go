package models

import "github.com/klearwave/service-info/pkg/models"

// VersionInput is the set of inputs required to create a version.
type VersionInput struct {
	VersionId string `json:"version_id,omitempty" example:"v0.1.2" doc:"Version in semantic versioning format."`

	ContainerImages []*ContainerImageInput `json:"container_images,omitempty" gorm:"many2many:version_container_images;" doc:"Container images associated with this version."`
}

// Version represents the full database schema for a Version.
type Version struct {
	models.Model
	VersionInput

	XVersion int  `json:"x_version,omitempty" example:"0" doc:"The X version of this release (e.g. 0; v0.1.2 == x.y.z)."`
	YVersion int  `json:"y_version,omitempty" example:"1" doc:"The Y version of this release (e.g. 1; v0.1.2 == x.y.z)."`
	ZVersion int  `json:"z_version,omitempty" example:"2" doc:"The Z version of this release (e.g. 2; v0.1.2 == x.y.z)."`
	Latest   bool `json:"latest,omitempty" example:"true" doc:"Whether or not this is the latest version."`
}

// VersionRequestCreate represents the request when creating a version.
type VersionRequestCreate struct {
	Body VersionInput
}

// VersionRequestGet represents the request when getting a version.
type VersionRequestGet struct {
	VersionID string `path:"version_id"`
}

// VersionRequestDelete represents the request when deleting a version.
type VersionRequestDelete struct {
	VersionID string `path:"version_id"`
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
