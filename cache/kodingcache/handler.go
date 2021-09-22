package kodingcache

import "samuelmanalu/go-cache-benchmark/cache"

func (k *kodingCacheInstance) Get(key string) (result cache.Data, err error) {
	cacheResult, err := k.cache.Get(key)
	result = cacheResult.(cache.Data)
	return
}

func (k *kodingCacheInstance) Set(key string, object cache.Data) (err error) {
	err = k.cache.Set(key, object)
	return
}
