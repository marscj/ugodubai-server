package service

import (
	"github.com/tiger1103/gfast-cache/cache"
)

type ICache interface {
	cache.IGCache
}

var c ICache

func Cache() ICache {
	if c == nil {
		panic("implement not found for interface ICache, forgot register?")
	}
	return c
}

func RegisterCache(che ICache) {
	c = che
}
