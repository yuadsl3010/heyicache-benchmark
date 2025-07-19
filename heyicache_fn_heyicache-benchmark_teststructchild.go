package main

import (
	"unsafe"

	"github.com/yuadsl3010/heyicache"
)

var (
	HeyiCacheFnStructSizeTestStructChild = int(unsafe.Sizeof(TestStructChild{}))
)

func HeyiCacheFnGetTestStructChild(bs []byte) interface{} {
	if len(bs) == 0 || len(bs) < HeyiCacheFnStructSizeTestStructChild {
		return nil
	}
	
	return (*TestStructChild)(unsafe.Pointer(&bs[0]))
}

func HeyiCacheFnSizeTestStructChild(value interface{}, isStructPtr bool) int32 {
	if value == nil {
		return 0
	}
	
	src, ok := value.(*TestStructChild)
	if !ok || src == nil {
		return 0
	}
	
	var size int32
	if isStructPtr {
		size = int32(HeyiCacheFnStructSizeTestStructChild)
	}
	// Id: success
	// TestName: success
	// string: foo string
	size += heyicache.HeyiCacheFnSizeString(src.TestName)
	// TestSkip: skip and set nil! struct tag skip
	// skip field: TestSkip
	return size
}

func HeyiCacheFnSetTestStructChild(value interface{}, bs []byte, isStructPtr bool) (interface{}, int32) {
	if value == nil {
		return nil, 0
	}
	
	src, ok := value.(*TestStructChild)
	if !ok || src == nil {
		return nil, 0
	}
	
	dst := src
	var size int32
	if isStructPtr {
		size = int32(HeyiCacheFnStructSizeTestStructChild)
		srcBytes := (*[1 << 30]byte)(unsafe.Pointer(src))[:size:size]
		copy(bs, srcBytes)
		dst = (*TestStructChild)(unsafe.Pointer(&bs[0]))
	}
	// Id: success
	// TestName: success
	// string: foo string
	pTestName, sizeTestName := heyicache.HeyiCacheFnSetString(src.TestName, bs[size:])
	size += sizeTestName
	dst.TestName = pTestName
	// TestSkip: skip and set nil! struct tag skip
	// skip field: TestSkip
	// dst.TestSkip = nil
	
	return dst, size
}

