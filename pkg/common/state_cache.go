package common

import (
	"github.com/patrickmn/go-cache"
	"time"
)

type StateCache struct {
	cache *cache.Cache
}

func NewStateCache() StateCache {
	return StateCache{
		cache: cache.New(5*time.Minute, 10*time.Minute),
	}
}

func (c *StateCache) Add(state string, ip string) {
	c.cache.Set(state, ip, cache.DefaultExpiration)
}

func (c *StateCache) Has(state string) bool {
	_, ok := c.cache.Get(state)
	return ok
}
