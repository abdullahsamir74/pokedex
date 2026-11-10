package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]CacheEntry
	mu    *sync.Mutex
}
type CacheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{
		cache: make(map[string]CacheEntry),
		mu:    &sync.Mutex{},
	}
	go c.reapLoop(interval)

	return &c

}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = CacheEntry{
		val:       value,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	cacheEntry, ok := c.cache[key]
	return cacheEntry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	timeAgo := time.Now().UTC().Add(-interval)
	for key, entry := range c.cache {
		if entry.createdAt.Before(timeAgo) {
			delete(c.cache, key)
		}
	}
}
