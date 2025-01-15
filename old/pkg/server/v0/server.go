package server

import (
	_ "net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/kataras/iris/v12"

	"github.com/klearwave/service-info/pkg/api/v0"
)

type Server struct{}

func NewServer() Server {
	return Server{}
}

func (s Server) CreateVersion(c *gin.Context) {

}

func (s Server) ReadVersion(c *gin.Context, versionId api.VersionId) {

}

func (s Server) ReadVersions(c *gin.Context) {

}

func (s Server) DeleteVersion(c *gin.Context, versionId api.VersionId) {

}
