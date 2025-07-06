package pokecache

import "time"

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entries: make(map[string]cacheEntry),
	}
	timeInterval := interval // time.Duration(interval) * time.Second
	go cache.reapLoop(timeInterval)
	return cache
}
