package main

import (
	"context"
	"time"

	"github.com/allegro/bigcache/v3"
)

// TestBigCache 使用 bigcache 包实现的 TestCacheIfc 接口
type TestBigCache struct {
	cache *bigcache.BigCache
}

// NewTestBigCache 创建一个新的 TestBigCache 实例
func NewTestBigCache(eviction time.Duration) (*TestBigCache, error) {
	config := bigcache.DefaultConfig(eviction)
	config.Verbose = false // 禁用日志输出
	cache, err := bigcache.New(context.Background(), config)
	if err != nil {
		return nil, err
	}

	return &TestBigCache{
		cache: cache,
	}, nil
}

// Get 实现 TestCacheIfc.Get 方法
func (b *TestBigCache) Get(key string) (*TestStruct, bool) {
	data, err := b.cache.Get(key)
	if err != nil {
		return nil, false
	}

	// 使用 protobuf 反序列化
	value, err := DeserializeTestStruct(data)
	if err != nil {
		return nil, false
	}

	return value, true
}

// Set 实现 TestCacheIfc.Set 方法
func (b *TestBigCache) Set(key string, value *TestStruct) error {
	// 使用 protobuf 序列化
	data, err := SerializeTestStruct(value)
	if err != nil {
		return err
	}

	return b.cache.Set(key, data)
}
