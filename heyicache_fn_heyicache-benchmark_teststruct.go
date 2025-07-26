package main

import (
	"unsafe"

	"github.com/yuadsl3010/heyicache"
)

var (
	// pass this ifc_ to heyicache in Get/Set
	HeyiCacheFnTestStructIfc_ = &HeyiCacheFnTestStructIfc{
		StructSize: int(unsafe.Sizeof(TestStruct{})),
	}
)

type HeyiCacheFnTestStructIfc struct {
	StructSize int
}

func (ifc *HeyiCacheFnTestStructIfc) Get (bs []byte) interface{} {
	if len(bs) == 0 || len(bs) < ifc.StructSize {
		return nil
	}
	
	return (*TestStruct)(unsafe.Pointer(&bs[0]))
}

func (ifc *HeyiCacheFnTestStructIfc) Size (value interface{}, isStructPtr bool) int32 {
	if value == nil {
		return 0
	}
	
	src, ok := value.(*TestStruct)
	if !ok || src == nil {
		return 0
	}
	
	var size int32
	if isStructPtr {
		size = int32(ifc.StructSize)
	}
	// Id: success
	// TestName: success
	// string: foo string
	size += heyicache.HeyiCacheFnSizeString(src.TestName)
	// TestSkip: success
	// string: foo string
	size += heyicache.HeyiCacheFnSizeString(src.TestSkip)
	// TestChild: success
	// struct: foo Foo
	size += HeyiCacheFnTestStructChildIfc_.Size(&src.TestChild, false)
	// TestChildren: success
	// slice struct: foo []Foo
	size += heyicache.HeyiCacheFnSizeSlice(src.TestChildren, HeyiCacheFnTestStructChildIfc_.StructSize)
	for idx := range src.TestChildren {
		size += HeyiCacheFnTestStructChildIfc_.Size(&src.TestChildren[idx], false)
	}
	// TestChildPtr: success
	// struct ptr: foo *Foo
	size += HeyiCacheFnTestStructChildIfc_.Size(src.TestChildPtr, true)
	// TestChildrenPtr: success
	// slice struct ptr: foo []*Foo
	size += heyicache.HeyiCacheFnSizeSlice(src.TestChildrenPtr, heyicache.HeyiCacheFnStructSizeptr)
	for _, item := range src.TestChildrenPtr {
		size += HeyiCacheFnTestStructChildIfc_.Size(item, true)
	}
	// TestProto: success
	// struct ptr: foo *Foo
	size += HeyiCacheFnTestPBIfc_.Size(src.TestProto, true)
	// Flag: success
	return size
}

func (ifc *HeyiCacheFnTestStructIfc) Set (value interface{}, bs []byte, isStructPtr bool) (interface{}, int32) {
	if value == nil {
		return nil, 0
	}
	
	src, ok := value.(*TestStruct)
	if !ok || src == nil {
		return nil, 0
	}
	
	dst := src
	var size int32
	if isStructPtr {
		size = int32(ifc.StructSize)
		srcBytes := (*[1 << 30]byte)(unsafe.Pointer(src))[:size:size]
		copy(bs, srcBytes)
		dst = (*TestStruct)(unsafe.Pointer(&bs[0]))
	}
	// Id: success
	// TestName: success
	// string: foo string
	pTestName, sizeTestName := heyicache.HeyiCacheFnSetString(src.TestName, bs[size:])
	size += sizeTestName
	dst.TestName = pTestName
	// TestSkip: success
	// string: foo string
	pTestSkip, sizeTestSkip := heyicache.HeyiCacheFnSetString(src.TestSkip, bs[size:])
	size += sizeTestSkip
	dst.TestSkip = pTestSkip
	// TestChild: success
	// struct: foo Foo
	_, sizeTestChild := HeyiCacheFnTestStructChildIfc_.Set(&dst.TestChild, bs[size:], false)
	size += sizeTestChild
	// TestChildren: success
	// slice struct: foo []Foo
	pTestChildren, sizeTestChildren := heyicache.HeyiCacheFnSetSlice(src.TestChildren, bs[size:], HeyiCacheFnTestStructChildIfc_.StructSize)
	size += sizeTestChildren
	dst.TestChildren = pTestChildren
	for idx := range dst.TestChildren {
		_, sizeTestChildren := HeyiCacheFnTestStructChildIfc_.Set(&dst.TestChildren[idx], bs[size:], false)
		size += sizeTestChildren
	}
	// TestChildPtr: success
	// struct ptr: foo *Foo
	pTestChildPtr, sizeTestChildPtr := HeyiCacheFnTestStructChildIfc_.Set(src.TestChildPtr, bs[size:], true)
	size += sizeTestChildPtr
	if pTestChildPtr != nil && sizeTestChildPtr > 0 {
		dst.TestChildPtr = pTestChildPtr.(*TestStructChild)
	}
	
	// TestChildrenPtr: success
	// slice struct ptr: foo []*Foo
	pTestChildrenPtr, sizeTestChildrenPtr := heyicache.HeyiCacheFnSetSlice(src.TestChildrenPtr, bs[size:], heyicache.HeyiCacheFnStructSizeptr)
	size += sizeTestChildrenPtr
	dst.TestChildrenPtr = pTestChildrenPtr
	for idx, item := range src.TestChildrenPtr {
		pTestChildrenPtr, sizeTestChildrenPtr := HeyiCacheFnTestStructChildIfc_.Set(item, bs[size:], true)
		size += sizeTestChildrenPtr
		if pTestChildrenPtr != nil && sizeTestChildrenPtr > 0 {
			dst.TestChildrenPtr[idx] = pTestChildrenPtr.(*TestStructChild)
		}
		
	}
	// TestProto: success
	// struct ptr: foo *Foo
	pTestProto, sizeTestProto := HeyiCacheFnTestPBIfc_.Set(src.TestProto, bs[size:], true)
	size += sizeTestProto
	if pTestProto != nil && sizeTestProto > 0 {
		dst.TestProto = pTestProto.(*TestPB)
	}
	
	// Flag: success
	
	return dst, size
}

