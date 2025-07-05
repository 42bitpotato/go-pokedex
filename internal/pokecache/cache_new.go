package pokecache

import "time"

func NewCache(interval int) {
	var cache Cache
	timeInterval := time.Duration(interval) * time.Second
	cache.reapLoop(timeInterval)
}
