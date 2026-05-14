package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	interval time.Duration

	mu      sync.Mutex
	entries map[string]cacheEntry
}

func NewCache(uInterval time.Duration) *Cache {
	c := &Cache{
		interval: uInterval,
		entries:  make(map[string]cacheEntry),
	}
}
