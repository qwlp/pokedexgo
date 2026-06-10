package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	items map[string]cacheEntry
	mu    sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		items: make(map[string]cacheEntry),
	}

	go c.readLoop(interval)

	return c
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	res, exists := c.items[key]
	if !exists {
		return nil, false
	}
	return res.val, true
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) readLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		now := time.Now()
		for k, i := range c.items {
			if now.Sub(i.createdAt) > interval {
				delete(c.items, k)
			}
		}
		c.mu.Unlock()
	}
}
