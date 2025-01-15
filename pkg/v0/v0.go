package v0

import "fmt"

const (
	DefaultPath = "/api/v0"
)

func PathFor(path string) string {
	return fmt.Sprintf("%s/%s", DefaultPath, path)
}
