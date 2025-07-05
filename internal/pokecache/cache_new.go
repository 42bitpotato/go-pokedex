package pokecache

import "time"

func NewCache(interval int) *Cache {
	cache := &Cache{
		entries: make(map[string]cacheEntry),
	}
	timeInterval := time.Duration(interval) * time.Second
	go cache.reapLoop(timeInterval)
	return cache
}
