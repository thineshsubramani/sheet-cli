package cache

import (
	"sync"
	"time"

	"github.com/thineshsubramani/sheet-cli/internal/types"
)

var (
	memCache      = make(map[string][]types.Section)
	memCacheTime  = make(map[string]time.Time)
	memCacheMutex sync.RWMutex
)

// getMemCache retrieves sections from the in-memory cache if valid.
func getMemCache(dbPath string, refreshCache bool) ([]types.Section, bool) {
	memCacheMutex.RLock()
	defer memCacheMutex.RUnlock()

	sections, ok := memCache[dbPath]
	cacheTime, hasTime := memCacheTime[dbPath]
	if ok && hasTime && time.Since(cacheTime) < cacheTTL && !refreshCache {
		return sections, true
	}
	return nil, false
}

// setMemCache updates the in-memory cache with new sections.
func setMemCache(dbPath string, sections []types.Section) {
	memCacheMutex.Lock()
	defer memCacheMutex.Unlock()

	memCache[dbPath] = sections
	memCacheTime[dbPath] = time.Now()
}
