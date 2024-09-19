package pokecache

import (
	"time"
)


type Cache struct {
	cache map[string]cacheEntry
}

type cacheEntry struct {
	value []byte
	createdAt time.Time
		// This allows us to say "after x time deelte cache"
}



func NewCache (interval time.Duration) (*Cache) {
	c := &Cache{
		cache: make(map[string]cacheEntry),
	}
	go c.reapLoop(interval)
	return c
}


func (c *Cache) Add(key string, val []byte) {
	c.cache[key] = cacheEntry {
		value: val,
		createdAt: time.Now().UTC(),
	}
}


func (c *Cache) Get(key string) ([]byte, bool) {
	cacheEnt, ok := c.cache[key]
	if !ok {
		return nil, false
	}
	return cacheEnt.value, true
}	



// Together, those "every 5 minutes (interval) delete any entries older than 5 minutes (interval)"

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		// This will run every interval miinutes  
		c.reap(interval)
	}
}


func (c *Cache) reap(interval time.Duration) {
	timeAgo := time.Now().UTC().Add(-interval)

	for key, val := range c.cache {
		if val.createdAt.Before(timeAgo) {
			delete(c.cache, key)
		}
	}
}




