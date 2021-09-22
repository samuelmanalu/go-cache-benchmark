package ristretto

import (
	"samuelmanalu/go-cache-benchmark/cache"

	"github.com/dgraph-io/ristretto"
)

type ristrettoCache struct {
	cache *ristretto.Cache
}

func New(size int) cache.Cache {
	memCache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: int64(size * 10), // number of keys to track frequency of (10M).
		MaxCost:     int64(size),      // maximum cost of cache (1GB).
		BufferItems: 64,               // number of keys per Get buffer.
		Metrics:     false,            // disable metrics for maximum performance
	})
	if err != nil {
		return nil
	}
	return &ristrettoCache{
		cache: memCache,
	}
}
