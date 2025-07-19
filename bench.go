package main

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/yuadsl3010/heyicache"
)

var (
	maxNum       = 1000000
	checkNum     = 100 // 1 write, 99 read times, the 100th time will check the data
	goroutineNum = 100 // 100 goroutines
)

type TestCacheIfc interface {
	Get(key string) (*TestStruct, bool)
	Set(key string, value *TestStruct) error
}

type BenchResult struct {
	ReadSuccess  uint64
	ReadMiss     uint64
	WriteSuccess uint64
	WriteFail    uint64
	CheckSuccess uint64
	CheckFail    uint64
}

func (result *BenchResult) String() string {
	readTotal := result.ReadSuccess + result.ReadMiss
	writeTotal := result.WriteSuccess + result.WriteFail
	checkTotal := result.CheckSuccess + result.CheckFail

	readFailRate := 0.0
	if readTotal > 0 {
		readFailRate = float64(result.ReadMiss) / float64(readTotal) * 100
	}
	writeFailRate := 0.0
	if writeTotal > 0 {
		writeFailRate = float64(result.WriteFail) / float64(writeTotal) * 100
	}
	checkFailRate := 0.0
	if checkTotal > 0 {
		checkFailRate = float64(result.CheckFail) / float64(checkTotal) * 100
	}

	return fmt.Sprintf(
		"\nRead: success=%d miss=%d missRate=%.2f%%\nWrite: success=%d fail=%d failRate=%.2f%%\nCheck: success=%d fail=%d failRate=%.2f%%",
		result.ReadSuccess, result.ReadMiss, readFailRate,
		result.WriteSuccess, result.WriteFail, writeFailRate,
		result.CheckSuccess, result.CheckFail, checkFailRate,
	)
}

func BenchIfc(b *testing.B, ifc TestCacheIfc) {
	result := &BenchResult{}
	wg := &sync.WaitGroup{}
	wg.Add(goroutineNum)
	for g := 0; g < goroutineNum; g++ {
		go func() {
			defer wg.Done()
			for i := 0; i < b.N; i++ {
				id := i % maxNum
				for j := 0; j < checkNum; j++ {
					if j%checkNum == 0 {
						// 1th set
						k, v := NewTestStruct(id)
						if err := ifc.Set(k, v); err != nil {
							result.WriteFail++
						} else {
							result.WriteSuccess++
						}
					} else {
						// 2~100th get
						v, ok := ifc.Get(GetKey(id))
						if ok {
							result.ReadSuccess++
							if j%checkNum == checkNum-1 {
								// 100th check
								if CheckTestStruct(id, v, false) {
									result.CheckSuccess++
								} else {
									result.CheckFail++
								}
							}
						} else {
							result.ReadMiss++
						}
					}
				}
			}
		}()
	}
	wg.Wait()
}

func BenchIfcForFreeCacheAndBigCache(b *testing.B, ifc TestCacheIfc) {
	result := &BenchResult{}
	wg := &sync.WaitGroup{}
	wg.Add(goroutineNum)
	for g := 0; g < goroutineNum; g++ {
		go func() {
			defer wg.Done()
			for i := 0; i < b.N; i++ {
				id := i % maxNum
				for j := 0; j < checkNum; j++ {
					if j%checkNum == 0 {
						// 1th set
						k, v := NewTestStruct(id)
						if err := ifc.Set(k, v); err != nil {
							result.WriteFail++
						} else {
							result.WriteSuccess++
						}
					} else {
						// 2~100th get
						v, ok := ifc.Get(GetKey(id))
						if ok {
							result.ReadSuccess++
							if j%checkNum == checkNum-1 {
								// 100th check
								if CheckTestStruct(id, v, true) {
									result.CheckSuccess++
								} else {
									result.CheckFail++
								}
							}
						} else {
							result.ReadMiss++
						}
					}
				}
			}
		}()
	}
	wg.Wait()
}

// only add lease logic for BenchIfc
func BenchHeyiCache(b *testing.B, heyi *TestHeyiCache) {
	result := &BenchResult{}
	wg := &sync.WaitGroup{}
	goroutineNum = 100
	wg.Add(goroutineNum)
	for g := 0; g < goroutineNum; g++ {
		go func() {
			defer wg.Done()
			for i := 0; i < b.N; i++ {
				id := i % maxNum
				for j := 0; j < checkNum; j++ {
					if j%checkNum == 0 {
						// 1th set
						k, v := NewTestStruct(id)
						if err := heyi.Set(k, v); err != nil {
							result.WriteFail++
						} else {
							result.WriteSuccess++
						}
					} else {
						ctx := heyicache.NewLeaseCtx(context.Background())
						leaseCtx := heyicache.GetLeaseCtx(ctx)
						leaseStoreCache := leaseCtx.GetLease(heyi.Cache)
						// 2~100th get
						v, ok := heyi.Get(leaseStoreCache, GetKey(id))
						if ok {
							result.ReadSuccess++
							if j%checkNum == checkNum-1 {
								// 100th check
								if CheckTestStruct(id, v, false) {
									result.CheckSuccess++
								} else {
									result.CheckFail++
								}
							}
						} else {
							result.ReadMiss++
						}

						heyicache.GetLeaseCtx(ctx).Done()
					}
				}
			}
		}()
	}
	wg.Wait()
	fmt.Println(result.String())
}
