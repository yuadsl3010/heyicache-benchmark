package main

import (
	"time"

	"github.com/patrickmn/go-cache"
)

// TestGoCache 使用 go-cache 包实现的 TestCacheIfc 接口
type TestGoCache struct {
	cache *cache.Cache
}

// NewTestGoCache 创建一个新的 TestGoCache 实例
func NewTestGoCache(defaultExpiration, cleanupInterval time.Duration) *TestGoCache {
	return &TestGoCache{
		cache: cache.New(defaultExpiration, cleanupInterval),
	}
}

// Get 实现 TestCacheIfc.Get 方法
func (g *TestGoCache) Get(key string) (*TestStruct, bool) {
	item, found := g.cache.Get(key)
	if !found {
		return nil, false
	}

	// 类型断言，将接口类型转换为 *TestStruct
	if value, ok := item.(*TestStruct); ok {
		return value, true
	}

	// 如果类型断言失败，返回 false
	return nil, false
}

// Set 实现 TestCacheIfc.Set 方法
func (g *TestGoCache) Set(key string, value *TestStruct) error {
	// 使用默认过期时间
	g.cache.Set(key, value, cache.DefaultExpiration)
	return nil
}
