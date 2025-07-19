package main

import (
	"github.com/coocood/freecache"
)

// TestFreeCache 使用 freecache 包实现的 TestCacheIfc 接口
type TestFreeCache struct {
	cache *freecache.Cache
}

// NewTestFreeCache 创建一个新的 TestFreeCache 实例
func NewTestFreeCache(cacheSize int) *TestFreeCache {
	return &TestFreeCache{
		cache: freecache.NewCache(cacheSize),
	}
}

// Get 实现 TestCacheIfc.Get 方法
func (f *TestFreeCache) Get(key string) (*TestStruct, bool) {
	data, err := f.cache.Get(StringToByte(key))
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
func (f *TestFreeCache) Set(key string, value *TestStruct) error {
	// 使用 protobuf 序列化
	data, err := SerializeTestStruct(value)
	if err != nil {
		return err
	}

	// freecache 需要指定过期时间（秒），这里设置为0表示永不过期
	return f.cache.Set(StringToByte(key), data, 0)
}
