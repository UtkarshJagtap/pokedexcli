package pokecache

import (
	"sync"
	"time"
)


type Cache struct {
	CacheEntries map[string]cacheEntry
  interval time.Duration
	mu           *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.CacheEntries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (cache *Cache) Get(key string) ([]byte, bool){
  cache.mu.Lock()
  defer cache.mu.Unlock()
  data, ok := cache.CacheEntries[key]
  if !ok {
    return nil, false
  }
  return data.val, ok

}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		CacheEntries: map[string]cacheEntry{},
    interval: interval,
    mu: &sync.Mutex{},
	}
	go cache.reapLoop()
	return cache
}

func (cache *Cache) reapLoop() {
  ticker := time.NewTicker(cache.interval)
  defer ticker.Stop()

  for range ticker.C{
    cache.mu.Lock()
    for key, entries := range cache.CacheEntries{
      if time.Since(entries.createdAt) >cache.interval{
        delete(cache.CacheEntries, key)
      } 
    }
    cache.mu.Unlock()
  }
}


