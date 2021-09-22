package freecache

import (
	"samuelmanalu/go-cache-benchmark/cache"

	"github.com/coocood/freecache"
)

type freeCacheInstance struct {
	cache *freecache.Cache
}

func New(size int) cache.Cache {
	freeCache := freecache.NewCache(size)
	return &freeCacheInstance{
		cache: freeCache,
	}
}
