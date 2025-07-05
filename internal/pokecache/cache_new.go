package pokecache

import "time"

func NewCache(interval time.Duration) {
	var cache Cache
	cache.reapLoop()
}
