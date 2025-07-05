package pokecache

import "time"

func NewCache(interval int) *Cache {
	cache := &Cache{}
	timeInterval := time.Duration(interval) * time.Second
	cache.reapLoop(timeInterval)
	return cache
}
