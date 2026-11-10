package pokecache

import (
	"sync"
	"time"
)

// Cache is a thread-safe, in-memory cache of byte slices. Entries older than
// the interval passed to NewCache are removed by a background goroutine.
type Cache struct {
	mu    sync.Mutex
	cache map[string]CacheEntry
}

type CacheEntry struct {
	val       []byte
	createdAt time.Time
}

// NewCache returns an empty cache that reaps entries older than interval.
func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		cache: make(map[string]CacheEntry),
	}
	go c.reapLoop(interval)

	return c
}

// Add stores value under key, replacing any existing entry.
func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = CacheEntry{
		val:       value,
		createdAt: time.Now().UTC(),
	}
}

// Get returns the value stored under key and whether the key was present.
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.cache[key]
	return entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

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
