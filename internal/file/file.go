package file

import (
	"fmt"
	"os"
	"path/filepath"
)

const cacheDir = ".sheet_cache"

// getCachePath returns the cache directory path.
func getCachePath() string {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "./" + cacheDir
	}
	return filepath.Join(dir, cacheDir)
}

// ensureCacheDir creates the cache directory if it doesn't exist.
func ensureCacheDir() error {
	return os.MkdirAll(getCachePath(), 0755)
}

// LocalCacheFile returns the full path to a cache file.
func LocalCacheFile(name string) string {
	return filepath.Join(getCachePath(), name)
}

// FileExists checks if a file exists.
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// ReadLocalFile reads a file from the local cache.
func ReadLocalFile(name string) ([]byte, error) {
	return os.ReadFile(LocalCacheFile(name))
}

// WriteLocalFile writes content to a local cache file.
func WriteLocalFile(name string, content []byte) error {
	if err := ensureCacheDir(); err != nil {
		return fmt.Errorf("failed to create cache directory: %w", err)
	}
	return os.WriteFile(LocalCacheFile(name), content, 0644)
}
