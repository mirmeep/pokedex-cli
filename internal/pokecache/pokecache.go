package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	CacheData map[string]cacheEntry
	mu sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	Val []byte
}

func NewCache(interval time.Duration) *Cache {
	newCache := Cache{
		CacheData: make(map[string]cacheEntry),
		interval: interval,
	}
	go newCache.reapLoop()
	return &newCache
}

func (c *Cache) Add(key string, val []byte) {
	newEntry := cacheEntry{
		createdAt: time.Now(),
		Val: val,
	}
	c.mu.Lock()
	c.CacheData[key] = newEntry
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	cacheVal, exists := c.CacheData[key]
	if exists {
		return cacheVal.Val, true
	} else {
		return nil, false
	}
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.mu.Lock()
		for key, data := range c.CacheData {
			if time.Since(data.createdAt) > c.interval {
				delete(c.CacheData, key)
			}
		}
		c.mu.Unlock()
	}	
}