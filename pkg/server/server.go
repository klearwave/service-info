package server

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"
	"github.com/gin-gonic/gin"

	"github.com/klearwave/service-info/pkg/v0/db"
	"github.com/klearwave/service-info/pkg/v0/models"
	"github.com/klearwave/service-info/pkg/v0/routes"
)

var version string = "unstable"

// server is a struct that represents a running server instance and parameters
// required to run a server instance.
type server struct {
	Router   *gin.Engine
	API      huma.API
	Database *db.Database
}

// NewServer returns a new instance of a server.  It returns an error if the
// server cannot be created.
func NewServer() (*server, error) {
	s := &server{Router: gin.Default()}

	config := huma.DefaultConfig("Information API", version)
	config.DocsPath = "/api/v0/docs"

	s.API = humagin.New(s.Router, config)

	db, err := db.NewDatabase()
	if err != nil {
		return nil, err
	}
	s.Database = db

	return s, nil
}

// RegisterRoutes registers all routes associated with a server instance.
func (s *server) RegisterRoutes() {
	// version routes
	version := models.Version{}
	huma.Register(s.API, routes.CreateVersion(&version), s.Database.CreateVersion)
	huma.Register(s.API, routes.GetVersion(&version), s.Database.GetVersion)
	huma.Register(s.API, routes.GetVersions(&version), s.Database.GetVersions)
	huma.Register(s.API, routes.DeleteVersion(&version), s.Database.DeleteVersion)
}
