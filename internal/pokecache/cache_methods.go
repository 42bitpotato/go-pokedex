package pokecache

import "time"

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.entries[key] = entry
}

func (c *Cache) Get(key string) (val []byte, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.entries[key]
	if !ok {
		return val, ok
	}
	val = entry.val
	return val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	for {
		time.Sleep(interval)
		if len(c.entries) > 0 {
			for key, entry := range c.entries {
				reapTime := entry.createdAt.Add(interval)
				if time.Now().After(reapTime) {
					delete(c.entries, key)
					continue
				}
			}
		}
		continue
	}
}
