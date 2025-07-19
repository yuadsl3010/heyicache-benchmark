module github.com/yuadsl3010/heyicache-benchmark

go 1.23

require (
	github.com/allegro/bigcache/v3 v3.1.0
	github.com/coocood/freecache v1.2.4
	github.com/gogo/protobuf v1.3.2
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/yuadsl3010/heyicache v0.0.0-20250719161959-09e61b70661e
)

require github.com/cespare/xxhash/v2 v2.3.0 // indirect

replace github.com/yuadsl3010/heyicache => ../heyicache
