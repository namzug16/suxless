package ui

import (
	"fmt"
)

// PublicFile generates a relative URL to a public file.
func PublicFile(filepath string) string {
	return fmt.Sprintf("/%s/%s", "files", filepath)
}

// StaticFile generates a relative URL to a static file including a cache-buster query parameter.
func StaticFile(filepath string) string {
	return fmt.Sprintf("/%s/%s", "static", filepath)
}
