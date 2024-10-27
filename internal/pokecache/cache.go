package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		mu:       &sync.Mutex{},
		interval: interval,
	}
	go cache.readLoop(interval)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.cache[key] = cacheEntry{
		createdAt: time.Time{},
		val:       val,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) (entryVal []byte, found bool) {
	c.mu.Lock()
	entry, found := c.cache[key]
	if found {
		entryVal = entry.val
	}
	c.mu.Unlock()
	return entryVal, found
}

func (c *Cache) readLoop(interval time.Duration) {
	ticker := time.Tick(interval)
	for range ticker {
		c.mu.Lock()
		for entry, entryVal := range c.cache {
			if time.Since(entryVal.createdAt) >= interval {
				delete(c.cache, entry)
			}
		}
		c.mu.Unlock()
	}
}
