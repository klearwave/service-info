package main

import "time"

// CreatedAt Creation timestamp of the version.  This is determined on the backend at creation time and is not modifiable.
type CreatedAt = time.Time

// Id Database ID of the stored object.  This is determined on the backend at creation time and is not modifiable.
type Id = int

// Latest Whether or not this version is the latest version.
type Latest = bool

// Version defines model for version.
type Version struct {
	// CreatedAt Creation timestamp of the version.  This is determined on the backend at creation time and is not modifiable.
	CreatedAt *CreatedAt `json:"created_at,omitempty" gorm:"primarykey"`

	// Id Database ID of the stored object.  This is determined on the backend at creation time and is not modifiable.
	Id *Id `json:"id,omitempty"`

	// Latest Whether or not this version is the latest version.
	Latest *Latest `json:"latest,omitempty"`

	// VersionId Version identifier in semantic versioning format.
	VersionId *VersionId `json:"version_id,omitempty"`
}

// VersionId Version identifier in semantic versioning format.
type VersionId = string
