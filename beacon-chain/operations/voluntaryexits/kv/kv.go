package kv

import (
	"github.com/patrickmn/go-cache"
)

// Cache defines the cache used to store voluntary exits.
type Cache struct {
	cache *cache.Cache
}

// NewCache initializes a cache for voluntary exits.
func NewCache() *Cache {
	pool := &Cache{
		cache: cache.New(cache.NoExpiration, 0 /* cleanupInterval */),
	}

	return pool
}
