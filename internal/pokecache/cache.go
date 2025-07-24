package pokecache

import ("sync"
		"time"
		"fmt")

type Cache struct {
	cacheMap map[string]CacheEntry
	mutex sync.Mutex
	interval time.Duration
}

type CacheEntry struct {
	createdAt time.Time
	val []byte
}

func reapLoop(cache *Cache) {
	// use the interval time to check if the a cacheEntry has expired, if yes
	// then remove the entry from the cacheMap
	ticker := time.NewTicker(cache.interval) 
	for {
		select {
		case <-ticker.C:   // go is very strict 
		 	cache.mutex.Lock()
			for key, value := range cache.cacheMap {
				if value.createdAt.Add(cache.interval).Before(time.Now()){
					delete(cache.cacheMap, key)
				}
			}
			cache.mutex.Unlock()
		}
	}
		
}

func NewCache(interval time.Duration) *Cache {
	cachePointer := &Cache{}
	cachePointer.cacheMap = make(map[string]CacheEntry)
	go reapLoop(cachePointer)
	return cachePointer
}

func (c Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	c.mutex.Unlock()
	newCache := CacheEntry{time.Now(), val}
	c.cacheMap[key] = newCache

}

func (c Cache) Get(key string) (val []byte, success bool){
	c.mutex.Lock()
	c.mutex.Unlock()
	value, exists := c.cacheMap[key]
	if exists {
		return value.val, true
	} else {
		return nil , false
	}
}