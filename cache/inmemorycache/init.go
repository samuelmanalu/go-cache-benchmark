package inmemorycache

import (
	"samuelmanalu/go-cache-benchmark/cache"
	"sync"
)

type inMemoryCache struct {
	cache map[string]cache.Data
	mutex *sync.RWMutex
}

func New() cache.Cache {
	return &inMemoryCache{
		cache: make(map[string]cache.Data),
		mutex: &sync.RWMutex{},
	}
}
