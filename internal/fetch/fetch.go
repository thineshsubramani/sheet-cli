package fetch

import (
	"io"
	"net/http"
	"strings"

	"github.com/thineshsubramani/sheet-cli/internal/file"
)

// FetchYAML fetches YAML content from a URL or local file.
func FetchYAML(path string) ([]byte, error) {
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		resp, err := http.Get(path)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		return io.ReadAll(resp.Body)
	}
	return file.ReadLocalFile(path)
}
