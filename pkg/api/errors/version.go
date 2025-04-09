package errors

import "errors"

var (
	ErrInvalidVersion = errors.New("invalid version object")

	ErrMissingVersionObject             = errors.New("missing version object")
	ErrMissingVersionParameterVersionId = errors.New("missing required version_id parameter for version")
	ErrMissingVersionParameterId        = errors.New("missing required id parameter for version")
)
