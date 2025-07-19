package main

import (
	"github.com/yuadsl3010/heyicache"
)

type TestHeyiCache struct {
	Cache *heyicache.Cache
}

// NewTestHeyiCache 创建一个新的 TestHeyiCache 实例
func NewTestHeyiCache(cacheSizeMB int) *TestHeyiCache {
	c, err := heyicache.NewCache(heyicache.Config{
		Name:    "TestHeyiCache",
		MaxSize: int32(cacheSizeMB),
	})
	if err != nil {
		panic(err)
	}
	return &TestHeyiCache{
		Cache: c,
	}
}

// Get 实现 TestCacheIfc.Get 方法
func (f *TestHeyiCache) Get(lease *heyicache.Lease, key string) (*TestStruct, bool) {
	data, err := f.Cache.Get(lease, StringToByte(key), HeyiCacheFnGetTestStruct)
	if err != nil || data == nil {
		return nil, false
	}

	return data.(*TestStruct), true
}

// Set 实现 TestCacheIfc.Set 方法
func (f *TestHeyiCache) Set(key string, value *TestStruct) error {
	return f.Cache.Set(StringToByte(key), value, HeyiCacheFnSetTestStruct, HeyiCacheFnSizeTestStruct, 0)
}

// func HeyiCacheFnGetTestStruct(data []byte) interface{} {
// 	return nil
// }

// func HeyiCacheFnSetTestStruct(value interface{}, bs []byte, _ bool) (interface{}, int32) {
// 	return nil, 0
// }

// func HeyiCacheFnSizeTestStruct(value interface{}, x bool) int32 {
// 	return 0
// }
