package models

import (
	"time"
)

// VersionBase is the base schema.
type VersionBase struct {
	Latest    bool   `json:"latest,omitempty" example:"true" doc:"Whether or not this is the latest version."`
	VersionId string `json:"version_id,omitempty" example:"v0.1.0" doc:"Version in semantic versioning format."`
}

// Version represents the full database schema for a Version.
type Version struct {
	VersionBase

	Id        int       `json:"id,omitempty" gorm:"primarykey" example:"1" doc:"Database ID of the stored object."`
	CreatedAt time.Time `json:"created_at,omitempty" example:"YYYY-MM-DDTHH:MM:SSZ" doc:"Object creation time in RFC 3339 format."`
}

// VersionRequestCreate represents the request when creating a version.
type VersionRequestCreate struct {
	Body VersionBase
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
