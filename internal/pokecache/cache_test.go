package pokecache

import (
    "testing"
    "time"
)

func TestAddGet(t *testing.T) {
    // Example: create a cache, add a key, get it, check expected value
	newCache := NewCache(time.Second * 3)
	newCache.Add("entry one", []byte("test"))
	if string(newCache.cacheMap["entry one"].val) != "test" {
		t.Error("value not matching")
	}

	data , success := newCache.Get("entry one")
	if success != true {
		t.Errorf("unable to find cache, found %d", data)
	}
}

