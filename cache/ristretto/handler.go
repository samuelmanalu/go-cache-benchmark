package ristretto

import (
	"errors"
	"samuelmanalu/go-cache-benchmark/cache"
)

func (r *ristrettoCache) Get(key string) (result cache.Data, err error) {
	cacheResult, isExist := r.cache.Get(key)
	if !isExist {
		err = errors.New("data not found")
		return
	}
	result = cacheResult.(cache.Data)
	return
}

func (r *ristrettoCache) Set(key string, object cache.Data) (err error) {
	isSuccessed := r.cache.Set(key, object, 1)
	if !isSuccessed {
		err = errors.New("store cache failed")
	}
	return
}
