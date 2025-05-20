package cache

import (
	"fmt"
	"time"

	"github.com/thineshsubramani/sheet-cli/internal/fetch"
	"github.com/thineshsubramani/sheet-cli/internal/file"
	"github.com/thineshsubramani/sheet-cli/internal/types"
	"gopkg.in/yaml.v3"
)

const cacheTTL = 10 * time.Minute

// RefreshAll refreshes the cache for all languages.
func RefreshAll(mainConfig *types.MainConfig, backend string) error {
	for _, lang := range mainConfig.Languages {
		fmt.Printf("Refreshing cache for language: %s\n", lang.Name)
		data, err := fetch.FetchYAML(backend + lang.Database)
		if err != nil {
			fmt.Printf("Failed to fetch %s: %v\n", lang.Database, err)
			continue
		}
		if err := file.WriteLocalFile(lang.Database, data); err != nil {
			fmt.Printf("Failed to write cache file %s: %v\n", lang.Database, err)
			continue
		}
		fmt.Printf("Cache updated: %s\n", lang.Database)
	}
	return nil
}

// LoadLanguageSections loads sections for a language from cache or backend.
func LoadLanguageSections(backend, dbPath string, refreshCache bool) ([]types.Section, error) {
	// Check memory cache
	if sections, ok := getMemCache(dbPath, refreshCache); ok {
		return sections, nil
	}

	// Check local cache file
	if !refreshCache && file.FileExists(file.LocalCacheFile(dbPath)) {
		data, err := file.ReadLocalFile(dbPath)
		if err == nil {
			var sections []types.Section
			if err := yaml.Unmarshal(data, &sections); err == nil {
				setMemCache(dbPath, sections)
				return sections, nil
			}
		}
	}

	// Fetch from backend
	data, err := fetch.FetchYAML(backend + dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch %s: %w", dbPath, err)
	}

	// Save to local cache
	if err := file.WriteLocalFile(dbPath, data); err != nil {
		fmt.Printf("Warning: failed to write cache file %s: %v\n", dbPath, err)
	}

	var sections []types.Section
	if err := yaml.Unmarshal(data, &sections); err != nil {
		return nil, fmt.Errorf("failed to unmarshal %s: %w", dbPath, err)
	}

	// Update memory cache
	setMemCache(dbPath, sections)
	return sections, nil
}
