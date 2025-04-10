package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	"github.com/klearwave/service-info/internal/pkg/api"
	apierrors "github.com/klearwave/service-info/internal/pkg/api/errors"
	"github.com/klearwave/service-info/internal/pkg/db"
)

// Service is an object that represents a service that interacts with a database.
type Service struct {
	Database *db.Database
}

func NewService(database *db.Database) *Service {
	return &Service{
		Database: database,
	}
}

// Create handles the creation of a new resource in the database for a web service.
func (service *Service) Create(ctx context.Context, req api.CreateRequest) *api.Result {
	return runRequest(service.Database, req, req.ToCreater().Create)
}

// Read handles the reading of an existing resource in the database for a web service.
func (service *Service) Read(ctx context.Context, req api.ReadRequest) *api.Result {
	return runRequest(service.Database, req, req.ToReader().Read)
}

// List handles the listing of existing resources in the database for a web service.
func (service *Service) List(ctx context.Context, req api.ListRequest) *api.Result {
	return runRequest(service.Database, req, req.ToLister().List)
}

// Update handles the updating of an existing resource in the database for a web service.
func (service *Service) Update(ctx context.Context, req api.UpdateRequest) *api.Result {
	return runRequest(service.Database, req, req.ToUpdater().Update)
}

// Delete handles the deleting of an existing resource from the database for a web service.
func (service *Service) Delete(ctx context.Context, req api.DeleteRequest) *api.Result {
	return runRequest(service.Database, req, req.ToDeleter().Delete)
}

type runRequestFunc func(*db.Database) *api.Result

// runRequest runs specific request logic.  It is used by top-level CRUD functions
// with the same logic applied to all.
func runRequest(database *db.Database, req api.Request, run runRequestFunc) *api.Result {
	database.Lock.Lock()
	defer database.Lock.Unlock()

	// ensure we have a valid database connection
	if database == nil {
		result := &api.Result{}
		result.InternalServerError(
			"database connection is invalid",
			fmt.Errorf("found a nil database connection"),
		)
	}

	// ensure we have a valid request function
	if run == nil {
		result := &api.Result{}
		result.InternalServerError(
			"request function is invalid",
			fmt.Errorf("found a nil request function"),
		)
	}

	// run pre-request validation
	if err := validateRequest(req); err != nil {
		return &api.Result{Error: err, Status: err.GetStatus()}
	}

	return run(database)
}

// validateRequest runs preflight operations on the request object.
func validateRequest(request api.Request) huma.StatusError {
	// ensure we have a request
	if request == nil {
		return apierrors.APIErrorFor(
			http.StatusBadRequest,
			"request is missing",
			nil,
		)
	}

	// authorize the request
	authorized, err := request.IsAuthorized()
	if err != nil || !authorized {
		return apierrors.APIErrorFor(
			http.StatusUnauthorized,
			"request is not authorized",
			err,
		)
	}

	// validate the request
	valid, err := request.IsValid()
	if err != nil || !valid {
		return apierrors.APIErrorFor(
			http.StatusBadRequest,
			"request is invalid",
			err,
		)
	}

	return nil
}
