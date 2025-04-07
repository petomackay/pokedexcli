package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	cache map[string]cacheEntry
	reapingInterval time.Duration 
	mu sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		cache:  make(map[string]cacheEntry),
		reapingInterval: interval,
	}

        go cache.reapLoop(interval)

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, present := c.cache[key]
	if !present {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		<- ticker.C
	        c.mu.Lock()
	        for key, entry := range c.cache {
	        	if time.Since(entry.createdAt) > interval {
	        		delete(c.cache, key)
	        	}
	        }
	        c.mu.Unlock()
	}
}
