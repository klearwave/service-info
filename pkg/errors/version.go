package errors

import "errors"

var (
	ErrInvalidVersion = errors.New("invalid version")

	ErrMissingVersionObject      = errors.New("missing version")
	ErrMissingVersionParameterId = errors.New("missing required version_id parameter for version")
)
