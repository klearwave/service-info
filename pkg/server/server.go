package server

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"
	"github.com/gin-gonic/gin"

	"github.com/klearwave/service-info/pkg/api"
	"github.com/klearwave/service-info/pkg/api/model/unversioned/route"
	routev0 "github.com/klearwave/service-info/pkg/api/model/v0/route"
	"github.com/klearwave/service-info/pkg/api/service"
	"github.com/klearwave/service-info/pkg/db"
)

// server is a struct that represents a running server instance and parameters
// required to run a server instance.
type server struct {
	Router   *gin.Engine
	API      huma.API
	Database *db.Database

	// services
	Service *service.Service
}

// NewServer returns a new instance of a server.  It returns an error if the
// server cannot be created.
func NewServer(connection *db.Connection) (*server, error) {
	s := &server{Router: gin.Default()}

	config := huma.DefaultConfig("Information API", api.ServerVersion)
	config.DocsPath = "/api/v0/docs"

	s.API = humagin.New(s.Router, config)

	db, err := db.NewDatabase(connection)
	if err != nil {
		return nil, err
	}
	s.Database = db
	s.Service = service.NewService(db)

	return s, nil
}

// RegisterRoutes registers all routes associated with a server instance.
func (s *server) RegisterRoutes() {
	// unversioned routes
	huma.Register(s.API, route.GetAbout(), s.Service.GetAbout)

	// version routes v0
	huma.Register(s.API, routev0.CreateVersion(), s.Service.CreateVersionV0)
	huma.Register(s.API, routev0.GetVersion(), s.Service.GetVersionV0)
	huma.Register(s.API, routev0.ListVersions(), s.Service.ListVersionsV0)
	huma.Register(s.API, routev0.ListVersionContainerImages(), s.Service.ListVersionContainerImagesV0)
	huma.Register(s.API, routev0.DeleteVersion(), s.Service.DeleteVersionV0)

	// container image routes v0
	huma.Register(s.API, routev0.CreateContainerImage(), s.Service.CreateContainerImageV0)
	huma.Register(s.API, routev0.GetContainerImage(), s.Service.GetContainerImageV0)
	huma.Register(s.API, routev0.ListContainerImages(), s.Service.ListContainerImagesV0)
	huma.Register(s.API, routev0.ListContainerImageVersions(), s.Service.ListContainerImageVersionsV0)
	huma.Register(s.API, routev0.DeleteContainerImage(), s.Service.DeleteContainerImageV0)
}
