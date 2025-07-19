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
