package unversioned

import (
	"github.com/klearwave/service-info/pkg/api"
	"github.com/klearwave/service-info/pkg/db"
)

type About struct {
	Version    string `json:"version,omitempty" example:"v0.1.2" doc:"Running version of this service in semantic versioning format."`
	CommitHash string `json:"commit_hash,omitempty" example:"631af50a8bbc4b5e69dab77d51a3a1733550fe8d" doc:"Commit hash of this running version."`
}

// Read handles the read request for an about model.
func (about *About) Read(database *db.Database) *api.Result {
	return &api.Result{
		Object: about,
		Error:  nil,
	}
}
