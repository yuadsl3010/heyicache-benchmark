package main

import (
	"unsafe"

	"github.com/yuadsl3010/heyicache"
)

var (
	HeyiCacheFnStructSizeTestPBChild = int(unsafe.Sizeof(TestPBChild{}))
)

func HeyiCacheFnGetTestPBChild(bs []byte) interface{} {
	if len(bs) == 0 || len(bs) < HeyiCacheFnStructSizeTestPBChild {
		return nil
	}
	
	return (*TestPBChild)(unsafe.Pointer(&bs[0]))
}

func HeyiCacheFnSizeTestPBChild(value interface{}, isStructPtr bool) int32 {
	if value == nil {
		return 0
	}
	
	src, ok := value.(*TestPBChild)
	if !ok || src == nil {
		return 0
	}
	
	var size int32
	if isStructPtr {
		size = int32(HeyiCacheFnStructSizeTestPBChild)
	}
	// Id: success
	// TestString: success
	// string: foo string
	size += heyicache.HeyiCacheFnSizeString(src.TestString)
	// TestStrings: success
	// slice string: foo []string
	size += heyicache.HeyiCacheFnSizeSlice(src.TestStrings, heyicache.HeyiCacheFnStructSizestring)
	for _, item := range src.TestStrings {
		size += heyicache.HeyiCacheFnSizeString(item)
	}
	// TestMap: skip and set nil! map type not supported cause it can't be stored by value, you must use custom serlization to store it if you really want map
	// skip field: TestMap
	// TestUint64S: success
	// slice: foo []int, []byte, etc.
	size += heyicache.HeyiCacheFnSizeSlice(src.TestUint64S, heyicache.HeyiCacheFnStructSizeuint64)
	// TestBytes: success
	// slice: foo []int, []byte, etc.
	size += heyicache.HeyiCacheFnSizeSlice(src.TestBytes, heyicache.HeyiCacheFnStructSizeuint8)
	// TestFloats: success
	// slice: foo []int, []byte, etc.
	size += heyicache.HeyiCacheFnSizeSlice(src.TestFloats, heyicache.HeyiCacheFnStructSizefloat32)
	return size
}

func HeyiCacheFnSetTestPBChild(value interface{}, bs []byte, isStructPtr bool) (interface{}, int32) {
	if value == nil {
		return nil, 0
	}
	
	src, ok := value.(*TestPBChild)
	if !ok || src == nil {
		return nil, 0
	}
	
	dst := src
	var size int32
	if isStructPtr {
		size = int32(HeyiCacheFnStructSizeTestPBChild)
		srcBytes := (*[1 << 30]byte)(unsafe.Pointer(src))[:size:size]
		copy(bs, srcBytes)
		dst = (*TestPBChild)(unsafe.Pointer(&bs[0]))
	}
	// Id: success
	// TestString: success
	// string: foo string
	pTestString, sizeTestString := heyicache.HeyiCacheFnSetString(src.TestString, bs[size:])
	size += sizeTestString
	dst.TestString = pTestString
	// TestStrings: success
	// slice string: foo []string
	pTestStrings, sizeTestStrings := heyicache.HeyiCacheFnSetSlice(src.TestStrings, bs[size:], heyicache.HeyiCacheFnStructSizestring)
	size += sizeTestStrings
	dst.TestStrings = pTestStrings
	for idx, item := range src.TestStrings {
		pTestStrings, sizeTestStrings := heyicache.HeyiCacheFnSetString(item, bs[size:])
		size += sizeTestStrings
		dst.TestStrings[idx] = pTestStrings
	}
	// TestMap: skip and set nil! map type not supported cause it can't be stored by value, you must use custom serlization to store it if you really want map
	// skip field: TestMap
	dst.TestMap = nil
	// TestUint64S: success
	// slice: foo []int, []byte, etc.
	pTestUint64S, sizeTestUint64S := heyicache.HeyiCacheFnSetSlice(src.TestUint64S, bs[size:], heyicache.HeyiCacheFnStructSizeuint64)
	size += sizeTestUint64S
	dst.TestUint64S = pTestUint64S
	// TestBytes: success
	// slice: foo []int, []byte, etc.
	pTestBytes, sizeTestBytes := heyicache.HeyiCacheFnSetSlice(src.TestBytes, bs[size:], heyicache.HeyiCacheFnStructSizeuint8)
	size += sizeTestBytes
	dst.TestBytes = pTestBytes
	// TestFloats: success
	// slice: foo []int, []byte, etc.
	pTestFloats, sizeTestFloats := heyicache.HeyiCacheFnSetSlice(src.TestFloats, bs[size:], heyicache.HeyiCacheFnStructSizefloat32)
	size += sizeTestFloats
	dst.TestFloats = pTestFloats
	
	return dst, size
}

