package main

import (
	"fmt"
	"testing"
	"time"
	"unsafe"
)

func BenchmarkMap(b *testing.B) {
	// init data
	ifc := &TestMap{
		c: make(map[string]*TestStruct, maxNum),
	}

	// run benchmark
	BenchIfc(b, ifc)
}

func BenchmarkGoCache(b *testing.B) {
	cache := NewTestGoCache(5*time.Minute, 10*time.Minute)
	BenchIfc(b, cache)
}

func BenchmarkFreeCache(b *testing.B) {
	// 设置缓存大小为100MB
	cache := NewTestFreeCache(100 * 1024 * 1024)
	BenchIfcForFreeCacheAndBigCache(b, cache)
}

func BenchmarkBigCache(b *testing.B) {
	// 设置过期时间为10分钟
	cache, err := NewTestBigCache(10 * time.Minute)
	if err != nil {
		b.Fatalf("Failed to create BigCache: %v", err)
	}
	BenchIfcForFreeCacheAndBigCache(b, cache)
}

func BenchmarkHeyiCache(b *testing.B) {
	// 设置缓存大小为100MB
	cache := NewTestHeyiCache(100)
	BenchHeyiCache(b, cache)
	totalEviction := cache.Cache.EvictionCount()         // 淘汰次数
	totalEvictionWait := cache.Cache.EvictionWaitCount() // 淘汰等待次数
	totalExpired := cache.Cache.ExpiredCount()           // 超时次数
	overwrites := cache.Cache.OverwriteCount()           // 覆盖次数
	hitCount := cache.Cache.HitCount()                   // 命中次数
	missCount := cache.Cache.MissCount()                 // 丢失次数
	lookupCount := cache.Cache.LookupCount()             // 命中 + 丢失
	hitRate := cache.Cache.HitRate()                     // 命中 / (命中 + 丢失)
	entryCount := cache.Cache.EntryCount()               // 总数
	fmt.Printf("totalEviction: %d\n", totalEviction)
	fmt.Printf("totalEvictionWait: %d\n", totalEvictionWait)
	fmt.Printf("totalExpired: %d\n", totalExpired)
	fmt.Printf("overwrites: %d\n", overwrites)
	fmt.Printf("hitCount: %d\n", hitCount)
	fmt.Printf("missCount: %d\n", missCount)
	fmt.Printf("lookupCount: %d\n", lookupCount)
	fmt.Printf("hitRate: %.2f\n", hitRate)
	fmt.Printf("entryCount: %d\n", entryCount)
}

func PrintString(testNamePtr *string) {
	fmt.Printf("str: %s\n", *testNamePtr)
	fmt.Printf("&str address: %p\n", testNamePtr)
	// string 在 Go 内部是一个结构体，包含指向底层数据的指针和长度
	stringHeader := (*struct {
		Data uintptr
		Len  int
	})(unsafe.Pointer(testNamePtr))
	fmt.Printf("str.Data address: %p\n", &stringHeader.Data)
	fmt.Printf("str.Len address: %p\n", &stringHeader.Len)
	fmt.Printf("str.Data -> address: 0x%x\n", stringHeader.Data)
	fmt.Println("")
}
