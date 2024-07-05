package cache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache map[string]cacheEntry
	mu    *sync.RWMutex
}

func CreateCache(clearInterval time.Duration) Cache {
	cache := Cache{
		cache: map[string]cacheEntry{},
		mu:    &sync.RWMutex{},
	}
	go cache.ReapLoop(clearInterval)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	return
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	data, ok := c.cache[key]
	if !ok {
		return nil, false
	}
	return data.val, true
}

func (c *Cache) ReapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mu.Lock()
		for key, val := range c.cache {
			if val.createdAt.Before(time.Now().Add(-interval)) {
				delete(c.cache, key)
			}
		}
		c.mu.Unlock()
	}
}
