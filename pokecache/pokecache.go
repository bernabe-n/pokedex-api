package pokecache

import "time"

type Cache struct {//Defines a struct named Cache
	cache map[string]cacheEntry //It has one field, cache: a map, key: string, value: cacheEntry
}

type cacheEntry struct {//Represents one item in the cache
	val			[]byte //the stored data (as bytes)
	createdAt	time.Time //timestamp when it was added
}

func NewCache(interval time.Duration) Cache { //Function that creates a new cache, Takes interval (how often cleanup runs), Returns a Cache value (not a pointer important).
	c := Cache{ //Creates a new Cache
		cache: make(map[string]cacheEntry), //Initializes the map using make
	}
	go c.reapLoop(interval) //Starts a goroutine that runs the reapLoop method, which will periodically clean up old entries based on the provided interval.
	return c
}

func (c *Cache) Add(key string, val []byte) {//Method to add an entry to the cache, takes a key and value (as bytes), c is a pointer receiver, meaning it modifies the Cache instance it is called on.
	c.cache[key] = cacheEntry{ //Adds a new entry to the cache map with the provided key and value. The value is wrapped in a cacheEntry struct that also records the current time as createdAt.
		val:       val,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) { //Method to retrieve an entry from the cache, takes a key, returns the value (as bytes) and a boolean indicating if it was found.
	cacheE, ok := c.cache[key]
	return cacheE.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) { //Method that runs in a loop to periodically clean up old entries from the cache, takes the interval as an argument.
	ticker := time.NewTicker(interval) //Creates a new ticker that will send a signal on its channel at regular intervals defined by the interval argument.
	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) { //Method to remove old entries from the cache, takes the interval as an argument.
	timeAgo := time.Now().UTC().Add(-interval)
	for k, v := range c.cache {
		if v.createdAt.Before(timeAgo) {
			delete(c.cache, k)
		}
	}
}