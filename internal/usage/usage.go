package usage

import (
	"encoding/json"
	"os"

	"github.com/thineshsubramani/sheet-cli/internal/file"
)

const usageFile = "usage.json"

// UsageMap tracks usage counts for languages.
type UsageMap map[string]int

// Track increments the usage count for a language.
func Track(lang string) {
	usage := UsageMap{}
	path := file.LocalCacheFile(usageFile)

	if content, err := os.ReadFile(path); err == nil {
		json.Unmarshal(content, &usage)
	}

	usage[lang]++
	data, _ := json.MarshalIndent(usage, "", "  ")
	os.WriteFile(path, data, 0644)
}
