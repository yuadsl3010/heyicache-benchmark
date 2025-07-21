package main

import (
	"testing"

	"github.com/yuadsl3010/heyicache"
)

func TestFnGenerateTool(t *testing.T) {
	heyicache.GenCacheFn(TestStruct{}, true)
}
