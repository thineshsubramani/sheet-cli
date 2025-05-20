package config

import (
	"fmt"
	"strings"

	"github.com/thineshsubramani/sheet-cli/internal/fetch"
	"github.com/thineshsubramani/sheet-cli/internal/file"
	"github.com/thineshsubramani/sheet-cli/internal/types"
	"gopkg.in/yaml.v3"
)

const mainConfigFile = "main.yaml"

// LoadMainConfig loads the main configuration from the backend or cache.
func LoadMainConfig(backend string, refreshCache bool) (*types.MainConfig, error) {
	cacheFile := mainConfigFile

	if refreshCache || !file.FileExists(file.LocalCacheFile(cacheFile)) {
		data, err := fetch.FetchYAML(backend + cacheFile)
		if err != nil {
			return nil, fmt.Errorf("fetch main.yaml failed: %w", err)
		}
		if err := file.WriteLocalFile(cacheFile, data); err != nil {
			return nil, fmt.Errorf("write main.yaml failed: %w", err)
		}
	}

	data, err := file.ReadLocalFile(cacheFile)
	if err != nil {
		return nil, fmt.Errorf("read main.yaml failed: %w", err)
	}

	var config types.MainConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("unmarshal main.yaml failed: %w", err)
	}
	return &config, nil
}

// FindLanguage finds a language by name or alias.
func FindLanguage(mainConfig *types.MainConfig, name string) *types.Language {
	for _, lang := range mainConfig.Languages {
		if strings.EqualFold(lang.Name, name) || contains(lang.Alias, name) {
			return &lang
		}
	}
	return nil
}

// ParseSectionInputs parses section inputs, handling spaces and commas.
func ParseSectionInputs(raw []string) []string {
	joined := strings.Join(raw, " ")
	joined = strings.ReplaceAll(joined, ",", " ")
	return strings.Fields(joined)
}

// contains checks if a string exists in a slice case-insensitively.
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if strings.EqualFold(s, item) {
			return true
		}
	}
	return false
}
