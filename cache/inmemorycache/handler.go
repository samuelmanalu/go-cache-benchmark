package inmemorycache

import (
	"errors"
	"samuelmanalu/go-cache-benchmark/cache"
)

func (i *inMemoryCache) Get(key string) (result cache.Data, err error) {
	i.mutex.RLock()
	result, isExist := i.cache[key]
	if !isExist {
		err = errors.New("data not found")
	}
	i.mutex.RUnlock()
	return
}

func (i *inMemoryCache) Set(key string, object cache.Data) (err error) {
	i.mutex.Lock()
	i.cache[key] = object
	i.mutex.Unlock()
	return
}
