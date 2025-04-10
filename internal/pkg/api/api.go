package api

import (
	"fmt"

	"github.com/klearwave/service-info/internal/pkg/db"
)

var (
	ServerVersion = "unstable"
	CommitHash    = "unknown"
)

const (
	DefaultPath = "/api"
)

type Reader interface {
	Read(*db.Database) *Result
}

type Lister interface {
	List(*db.Database) *Result
}

type Creater interface {
	Create(*db.Database) *Result
}

type Updater interface {
	Update(*db.Database) *Result
}

type Deleter interface {
	Delete(*db.Database) *Result
}

func PathFor(path string) string {
	return fmt.Sprintf("%s/%s", DefaultPath, path)
}
