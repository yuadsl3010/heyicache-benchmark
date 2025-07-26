package main

import (
	"fmt"
	"testing"
	"time"
	"unsafe"
)

func BenchmarkMap(b *testing.B) {
	// return
	// init data
	ifc := &TestMap{
		c: make(map[string]*TestStruct, maxNum),
	}

	// run benchmark
	BenchIfc(b, ifc)
}

func BenchmarkGoCache(b *testing.B) {
	// return
	cache := NewTestGoCache(5*time.Minute, 10*time.Minute)
	BenchIfc(b, cache)
}

func BenchmarkFreeCache(b *testing.B) {
	// return
	// 设置缓存大小为100MB
	cache := NewTestFreeCache(100 * 1024 * 1024)
	BenchIfcForFreeCacheAndBigCache(b, cache)
}

func BenchmarkBigCache(b *testing.B) {
	// return
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

	// evictionNum := cache.Cache.EvictionNum()             // 淘汰个数
	// evictionCount := cache.Cache.EvictionCount()         // 淘汰触发次数
	// evictionWaitCount := cache.Cache.EvictionWaitCount() // 淘汰次数
	// expireCount := cache.Cache.ExpireCount()             // 超时次数
	// hitCount := cache.Cache.HitCount()                   // 命中次数
	// missCount := cache.Cache.MissCount()                 // 丢失次数
	// readCount := cache.Cache.ReadCount()                 // 命中 + 丢失
	// writeCount := cache.Cache.WriteCount()               // 写入次数
	// writeErrCount := cache.Cache.WriteErrCount()         // 写入错误次数
	// overwriteCount := cache.Cache.OverwriteCount()       // 覆盖次数
	// skipWriteCount := cache.Cache.SkipWriteCount()       // 跳过写入次数
	// used, mem := cache.Cache.MemStat()
	// entryCount := cache.Cache.EntryCount() // 总数
	// fmt.Printf("evictionNum: %d\n", evictionNum)
	// fmt.Printf("evictionCount: %d\n", evictionCount)
	// fmt.Printf("evictionWaitCount: %d\n", evictionWaitCount)
	// fmt.Printf("expireCount: %d\n", expireCount)
	// fmt.Printf("hitCount: %d\n", hitCount)
	// fmt.Printf("missCount: %d\n", missCount)
	// fmt.Printf("readCount: %d\n", readCount)
	// fmt.Printf("writeCount: %d\n", writeCount)
	// fmt.Printf("writeErrCount: %d\n", writeErrCount)
	// fmt.Printf("overwriteCount: %d\n", overwriteCount)
	// fmt.Printf("skipWriteCount: %d\n", skipWriteCount)
	// fmt.Printf("used: %d\n", used)
	// fmt.Printf("mem: %d\n", mem)
	// fmt.Printf("entryCount: %d\n", entryCount)
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

// func TestYZCHeyiCache(b *testing.T) {
// 	// 设置缓存大小为100MB
// 	cache := NewTestHeyiCache(100)

// 	ctx := heyicache.NewLeaseCtx(context.Background())
// 	leaseCtx := heyicache.GetLeaseCtx(ctx)
// 	leaseStoreCache := leaseCtx.GetLease(cache.Cache)
// 	id := 100

// 	k, src := NewTestStruct(id)
// 	err := cache.Cache.Set(StringToByte(k), src, HeyiCacheFnTestStructIfc_, 0)
// 	if err != nil {
// 		panic(err)
// 	}

// 	r, err := cache.Cache.Get(leaseStoreCache, StringToByte(k), HeyiCacheFnTestStructIfc_)
// 	if err != nil {
// 		panic(err)
// 	}

// 	dst := r.(*TestStruct)
// 	fmt.Println("dst", dst)
// 	fmt.Println("diff", CheckTestStruct(id, dst, false))
// }
