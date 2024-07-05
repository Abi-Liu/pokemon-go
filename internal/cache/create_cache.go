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
	return Cache{
		cache: map[string]cacheEntry{},
		mu:    &sync.RWMutex{},
	}
}

func (c Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	return
}

func (c Cache) Get(key string) ([]byte, bool) {
	data, ok := c.cache[key]
	if !ok {
		return nil, false
	}
	return data.val, true
}
