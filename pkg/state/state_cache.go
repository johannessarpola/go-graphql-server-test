package state

import (
	"github.com/patrickmn/go-cache"
	"time"
)

type Cache struct {
	cache *cache.Cache
}

func NewStateCache() Cache {
	return Cache{
		cache: cache.New(5*time.Minute, 10*time.Minute),
	}
}

func (c *Cache) Add(state string, ip string) {
	c.cache.Set(state, ip, cache.DefaultExpiration)
}

func (c *Cache) Has(state string) bool {
	_, ok := c.cache.Get(state)
	return ok
}
