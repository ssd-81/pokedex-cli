package pokecache

import ("sync"
		"time"
		)

type Cache struct {
	cacheMap map[string]CacheEntry
	mutex sync.Mutex
	interval time.Duration
}

type CacheEntry struct {
	createdAt time.Time
	val []byte
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval) 
	for {
		select {
		case <-ticker.C:   // go is very strict 
		 	c.mutex.Lock()
			for key, value := range c.cacheMap {
				if value.createdAt.Add(c.interval).Before(time.Now()){
					delete(c.cacheMap, key)
				}
			}
			c.mutex.Unlock()
		}
	}	
}

func NewCache(interval time.Duration) *Cache {
	cachePointer := &Cache{}
	cachePointer.cacheMap = make(map[string]CacheEntry)
	go cachePointer.reapLoop()
	return cachePointer
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	newCache := CacheEntry{time.Now(), val}
	c.cacheMap[key] = newCache

}

func (c *Cache) Get(key string) (val []byte, success bool){
	c.mutex.Lock()
	defer c.mutex.Unlock()
	value, exists := c.cacheMap[key]
	if exists {
		return value.val, true
	} else {
		return nil , false
	}
}