package kodingcache

import (
	"time"

	koding "github.com/koding/cache"

	"samuelmanalu/go-cache-benchmark/cache"
)

type kodingCacheInstance struct {
	cache koding.Cache
}

func New(ttl time.Duration) cache.Cache {
	memCache := koding.NewMemoryWithTTL(ttl)
	return &kodingCacheInstance{
		cache: memCache,
	}
}
