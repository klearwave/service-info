package test

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"

	"github.com/klearwave/service-info/internal/pkg/api"
	"github.com/klearwave/service-info/internal/pkg/api/model"
	modelv0 "github.com/klearwave/service-info/internal/pkg/api/model/v0"
	v0 "github.com/klearwave/service-info/internal/pkg/api/model/v0"
	createv0 "github.com/klearwave/service-info/internal/pkg/api/model/v0/request/create"
	deletev0 "github.com/klearwave/service-info/internal/pkg/api/model/v0/request/delete"
	listv0 "github.com/klearwave/service-info/internal/pkg/api/model/v0/request/list"
	readv0 "github.com/klearwave/service-info/internal/pkg/api/model/v0/request/read"
	routev0 "github.com/klearwave/service-info/internal/pkg/api/model/v0/route"
	"github.com/klearwave/service-info/internal/pkg/db"
	"github.com/klearwave/service-info/internal/pkg/server"
	"github.com/klearwave/service-info/internal/pkg/utils/pointers"
)

var authHeader = map[string]string{
	"Authorization": "Basic " + base64.StdEncoding.EncodeToString([]byte(api.AuthUsername+":"+api.AuthPassword)),
}

// httpTests defines the httpTests to run against the REST API.  These will be executed in order
// and are not parallelized.
var httpTests = []struct {
	name           string
	request        interface{}            // request payload sent (to be converted to map[string]interface)
	response       map[string]interface{} // expected response (if any)
	headers        map[string]string
	httpPath       string
	method         string
	expectedStatus int
}{
	{
		name: "fail: unauthorized version request (create)",
		request: createv0.VersionBody{
			VersionBase: modelv0.VersionBase{
				Id: pointers.FromString("x.y.z"),
			},
		},
		httpPath:       v0.PathFor(routev0.DefaultVersionsPath),
		method:         http.MethodPost,
		expectedStatus: http.StatusUnauthorized,
	},
	{
		name: "fail: unauthorized container image request (create)",
		request: createv0.ContainerImageBody{
			ContainerImageBase: modelv0.ContainerImageBase{
				Image:      pointers.FromString("postgres2"),
				SHA256Sum:  pointers.FromString("2d4b92db6941294f731cfe7aeca336eb8dba279171c0e6ceda32b9f018f8429d"),
				CommitHash: pointers.FromString("631af50a8bbc4b5e69dab77d51a3a1733550fe8d"),
			},
		},
		httpPath:       v0.PathFor(routev0.DefaultContainerImagesPath),
		method:         http.MethodPost,
		expectedStatus: http.StatusUnauthorized,
	},
	{
		name: "fail: ensure invalid version",
		request: createv0.VersionBody{
			VersionBase: modelv0.VersionBase{
				Id: pointers.FromString("x.y.z"),
			},
		},
		httpPath:       v0.PathFor(routev0.DefaultVersionsPath),
		method:         http.MethodPost,
		headers:        authHeader,
		expectedStatus: http.StatusBadRequest,
	},
	{
		name: "fail: ensure invalid container image (missing sha256sum)",
		request: createv0.ContainerImageBody{
			ContainerImageBase: modelv0.ContainerImageBase{
				Image:      pointers.FromString("postgres"),
				CommitHash: pointers.FromString("631af50a8bbc4b5e69dab77d51a3a1733550fe8d"),
			},
		},
		httpPath:       v0.PathFor(routev0.DefaultContainerImagesPath),
		method:         http.MethodPost,
		headers:        authHeader,
		expectedStatus: http.StatusBadRequest,
	},
	{
		name: "fail: ensure invalid container image (missing commit hash)",
		request: createv0.ContainerImageBody{
			ContainerImageBase: modelv0.ContainerImageBase{
				Image:     pointers.FromString("postgres"),
				SHA256Sum: pointers.FromString("2d4b92db6941294f731cfe7aeca336eb8dba279171c0e6ceda32b9f018f8429d"),
			},
		},
		httpPath:       v0.PathFor(routev0.DefaultContainerImagesPath),
		method:         http.MethodPost,
		headers:        authHeader,
		expectedStatus: http.StatusBadRequest,
	},
	{
		name: "fail: ensure missing version is not found",
		request: readv0.VersionRequest{
			StringFetcher: model.StringFetcher{
				Id: "v0.0.1",
			},
		},
		httpPath:       v0.PathFor(routev0.DefaultVersionsPath) + "/v0.0.1",
		method:         http.MethodGet,
		expectedStatus: http.StatusNotFound,
	},
	{
		name: "fail: ensure missing container image is not found",
		request: readv0.ContainerImageRequest{
			IntegerFetcher: model.IntegerFetcher{
				Id: 999,
			},
		},
		httpPath:       v0.PathFor(routev0.DefaultContainerImagesPath) + "/999",
		method:         http.MethodGet,
		expectedStatus: http.StatusNotFound,
	},
	{
		name:           "success: ensure empty versions returns",
		request:        listv0.VersionRequest{},
		httpPath:       v0.PathFor(routev0.DefaultVersionsPath),
		method:         http.MethodGet,
		expectedStatus: http.StatusOK,
	},
	{
		name:           "success: ensure empty container images returns",
		request:        listv0.ContainerImageRequest{},
		httpPath:       v0.PathFor(routev0.DefaultContainerImagesPath),
		method:         http.MethodGet,
		expectedStatus: http.StatusOK,
	},
	{
		name: "success: ensure version is created successfully without container images",
		request: createv0.VersionBody{
			VersionBase: modelv0.VersionBase{
				Id: pointers.FromString("v0.1.2"),
			},
		},
		httpPath:       v0.PathFor(routev0.DefaultVersionsPath),
		method:         http.MethodPost,
		headers:        authHeader,
		expectedStatus: http.StatusOK,
	},
	{
		name: "fail: ensure version uniquness",
		request: createv0.VersionBody{
			VersionBase: modelv0.VersionBase{
				Id: pointers.FromString("v0.1.2"),
			},
		},
		httpPath:       v0.PathFor(routev0.DefaultVersionsPath),
		method:         http.MethodPost,
		headers:        authHeader,
		expectedStatus: http.StatusInternalServerError,
	},
	{
		name: "success: ensure version is created successfully with container images",
		request: createv0.VersionBody{
			VersionBase: modelv0.VersionBase{
				Id: pointers.FromString("v0.1.3"),
			},
			ContainerImages: []*createv0.ContainerImageBody{
				{
					ContainerImageBase: modelv0.ContainerImageBase{
						Image:      pointers.FromString("postgres"),
						SHA256Sum:  pointers.FromString("2d4b92db6941294f731cfe7aeca336eb8dba279171c0e6ceda32b9f018f8429d"),
						CommitHash: pointers.FromString("631af50a8bbc4b5e69dab77d51a3a1733550fe8d"),
					},
				},
			},
		},
		httpPath:       v0.PathFor(routev0.DefaultVersionsPath),
		method:         http.MethodPost,
		headers:        authHeader,
		expectedStatus: http.StatusOK,
	},
	{
		name: "success: container image exists",
		request: readv0.ContainerImageRequest{
			IntegerFetcher: model.IntegerFetcher{
				Id: 1,
			},
		},
		httpPath:       v0.PathFor(routev0.DefaultContainerImagesPath) + "/1",
		method:         http.MethodGet,
		expectedStatus: http.StatusOK,
	},
	{
		name: "fail: unauthorized version request (delete)",
		request: deletev0.VersionRequest{
			StringFetcher: model.StringFetcher{
				Id: "v0.1.3",
			},
		},
		httpPath:       v0.PathFor(routev0.DefaultVersionsPath) + "/v0.1.3",
		method:         http.MethodDelete,
		expectedStatus: http.StatusUnauthorized,
	},
	{
		name: "fail: unauthorized container image request (delete)",
		request: deletev0.ContainerImageRequest{
			IntegerFetcher: model.IntegerFetcher{
				Id: 1,
			},
		},
		httpPath:       v0.PathFor(routev0.DefaultContainerImagesPath) + "/1",
		method:         http.MethodDelete,
		expectedStatus: http.StatusUnauthorized,
	},
	{
		name: "success: ensure version is deleted successfully",
		request: deletev0.VersionRequest{
			StringFetcher: model.StringFetcher{
				Id: "v0.1.3",
			},
		},
		httpPath:       v0.PathFor(routev0.DefaultVersionsPath) + "/v0.1.3",
		method:         http.MethodDelete,
		headers:        authHeader,
		expectedStatus: http.StatusOK,
	},
	{
		name: "success: container image still exists",
		request: readv0.ContainerImageRequest{
			IntegerFetcher: model.IntegerFetcher{
				Id: 1,
			},
		},
		httpPath:       v0.PathFor(routev0.DefaultContainerImagesPath) + "/1",
		method:         http.MethodGet,
		expectedStatus: http.StatusOK,
	},
	{
		name: "success: create version with linked existing container image",
		request: createv0.VersionBody{
			VersionBase: modelv0.VersionBase{
				Id: pointers.FromString("v0.1.3"),
			},
			ContainerImages: []*createv0.ContainerImageBody{
				{
					ContainerImageBase: modelv0.ContainerImageBase{
						Image:      pointers.FromString("postgres"),
						SHA256Sum:  pointers.FromString("2d4b92db6941294f731cfe7aeca336eb8dba279171c0e6ceda32b9f018f8429d"),
						CommitHash: pointers.FromString("631af50a8bbc4b5e69dab77d51a3a1733550fe8d"),
					},
				},
			},
		},
		httpPath:       v0.PathFor(routev0.DefaultVersionsPath),
		method:         http.MethodPost,
		headers:        authHeader,
		expectedStatus: http.StatusOK,
	},
	{
		name: "success: create version with mismatched existing container image",
		request: createv0.VersionBody{
			VersionBase: modelv0.VersionBase{
				Id: pointers.FromString("v0.1.4"),
			},
			ContainerImages: []*createv0.ContainerImageBody{
				{
					ContainerImageBase: modelv0.ContainerImageBase{
						Image:      pointers.FromString("postgres"),
						SHA256Sum:  pointers.FromString("2d4b92db6941294f731cfe7aeca336eb8dba279171c0e6ceda32b9f018f8429e"),
						CommitHash: pointers.FromString("631af50a8bbc4b5e69dab77d51a3a1733550fe8e"),
					},
				},
			},
		},
		httpPath:       v0.PathFor(routev0.DefaultVersionsPath),
		method:         http.MethodPost,
		headers:        authHeader,
		expectedStatus: http.StatusInternalServerError,
	},
	{
		name: "success: create new container image",
		request: createv0.ContainerImageBody{
			ContainerImageBase: modelv0.ContainerImageBase{
				Image:      pointers.FromString("postgres2"),
				SHA256Sum:  pointers.FromString("2d4b92db6941294f731cfe7aeca336eb8dba279171c0e6ceda32b9f018f8429d"),
				CommitHash: pointers.FromString("631af50a8bbc4b5e69dab77d51a3a1733550fe8d"),
			},
		},
		httpPath:       v0.PathFor(routev0.DefaultContainerImagesPath),
		method:         http.MethodPost,
		headers:        authHeader,
		expectedStatus: http.StatusOK,
	},
	{
		name: "fail: create existing container image",
		request: createv0.ContainerImageBody{
			ContainerImageBase: modelv0.ContainerImageBase{
				Image:      pointers.FromString("postgres"),
				SHA256Sum:  pointers.FromString("2d4b92db6941294f731cfe7aeca336eb8dba279171c0e6ceda32b9f018f8429d"),
				CommitHash: pointers.FromString("631af50a8bbc4b5e69dab77d51a3a1733550fe8d"),
			},
		},
		httpPath:       v0.PathFor(routev0.DefaultContainerImagesPath),
		method:         http.MethodPost,
		headers:        authHeader,
		expectedStatus: http.StatusInternalServerError,
	},
	{
		name: "success: delete container image",
		request: deletev0.ContainerImageRequest{
			IntegerFetcher: model.IntegerFetcher{
				Id: 1,
			},
		},
		httpPath:       v0.PathFor(routev0.DefaultContainerImagesPath) + "/1",
		method:         http.MethodDelete,
		headers:        authHeader,
		expectedStatus: http.StatusOK,
	},
}

func TestE2E(t *testing.T) {
	connection := &db.Config{
		Host:         "localhost",
		Port:         5432,
		Username:     "postgres",
		Password:     "postgres",
		DatabaseName: "postgres",
	}

	// create the embedded postgres database for testing if requested
	if os.Getenv("E2E_EMBEDDED") == "true" {
		database := embeddedpostgres.NewDatabase()
		if err := database.Start(); err != nil {
			t.Fatal(err)
		}

		defer func() {
			if err := database.Stop(); err != nil {
				t.Fatal(err)
			}
		}()
	} else {
		if err := connection.Parse(); err != nil {
			t.Fatal(err)
		}
	}

	// create the in memory http server for testing if requested
	serverURL := "http://localhost:8888"

	if os.Getenv("E2E_EMBEDDED") == "true" {
		s, err := server.NewServer()
		if err != nil {
			t.Fatal(err)
		}
		s.RegisterRoutes()

		if err := s.Init(connection); err != nil {
			t.Fatal(err)
		}

		server := httptest.NewServer(s.Router)
		serverURL = server.URL
		defer server.Close()
	}

	// run migrations first if requested
	if os.Getenv("E2E_EMBEDDED") == "true" {
		migrate(t, connection)
	}

	client := &http.Client{}

	// run individual tests next
	for _, tc := range httpTests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			// NO t.Parallel() — ensures ordered execution
			body, err := json.Marshal(tc.request)
			if err != nil {
				t.Fatalf("Failed to marshal JSON: %v", err)
			}

			httpPath := serverURL + tc.httpPath

			req, err := http.NewRequest(tc.method, serverURL+tc.httpPath, bytes.NewBuffer(body))
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			// set headers
			req.Header.Set("Content-Type", "application/json")
			if tc.headers != nil {
				for k, v := range tc.headers {
					req.Header.Set(k, v)
				}
			}

			resp, err := client.Do(req)
			if err != nil {
				t.Fatalf("Failed to send request: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tc.expectedStatus {
				t.Errorf("httpPath [%s]: Expected status %d, got %d", httpPath, tc.expectedStatus, resp.StatusCode)
			}
		})
	}
}

// connect makes a database connection.
func connect(t *testing.T, connection *db.Config) *sqlx.DB {
	db, err := sqlx.Connect("postgres", connection.String)
	if err != nil {
		t.Fatal(err)
	}

	return db
}

// migrate runs database migrations.  This is to be used prior to executing any tests
// and ensures our migrations run correctly.
func migrate(t *testing.T, connection *db.Config) {
	db := connect(t, connection)

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal(fmt.Errorf("unable to get caller information"))
	}

	baseDir := filepath.Dir(filename)

	if err := goose.Up(db.DB, filepath.Join(baseDir, "../migrations")); err != nil {
		t.Fatal(err)
	}
}
