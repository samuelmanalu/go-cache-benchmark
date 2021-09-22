package freecache

import (
	"encoding/json"
	"samuelmanalu/go-cache-benchmark/cache"
)

func (f *freeCacheInstance) Get(key string) (result cache.Data, err error) {
	cacheResult, err := f.cache.Get([]byte(key))
	if err != nil {
		return
	}
	if err = json.Unmarshal(cacheResult, &result); err != nil {
		return
	}
	return
}

func (f *freeCacheInstance) Set(key string, object cache.Data) (err error) {
	data, err := json.Marshal(object)
	if err != nil {
		return
	}
	err = f.cache.Set([]byte(key), data, 0)
	return
}
