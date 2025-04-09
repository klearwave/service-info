package read

import (
	"github.com/klearwave/service-info/pkg/api"
	"github.com/klearwave/service-info/pkg/api/model"
	v0 "github.com/klearwave/service-info/pkg/api/model/v0"
)

// ContainerImageRequest represents the request when reading a specific container image.
type ContainerImageRequest struct {
	model.IntegerFetcher
}

// IsAuthorized checks if the request is authorized.  It is used to satisfy the Request.
func (req ContainerImageRequest) IsAuthorized() (bool, error) {
	return true, nil
}

// IsValid checks if the request is valid.  It is used to satisfy the Reader interface.
func (req ContainerImageRequest) IsValid() (bool, error) {
	return req.IntegerFetcher.IsValid()
}

// ToReader converts the request to a reader object.  It is used to satisfy the
// Reader interface.
func (req ContainerImageRequest) ToReader() api.Reader {
	return &v0.ContainerImage{
		ModelWithId: model.ModelWithId{
			Id: req.Id,
		},
	}
}
