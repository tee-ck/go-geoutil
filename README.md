# go-geoutil

A Go library for geolocation utilities. It provides a set of functions to measure distances.

# Installation

```shell
go get github.com/tee-ck/geoutil
```

# Usage

```go
package main

import (
	"fmt"
	"github.com/tee-ck/geoutil"
)

func main() {
	p1 := geoutil.NewPoint(3.096249444382203, 101.53713780926184)
	p2 := geoutil.NewPoint(1.3361333233860981, 103.83548216216576)
	
	fmt.Println(p1.DistanceTo(p2)) // 321.74 meters
}
```

# Benchmarks

```shell
cheekuan@linux:go-geoutil$ go test -bench . -run ^Bench -benchtime 10000000x -benchmem
goos: linux
goarch: amd64
pkg: github.com/tee-ck/go-geoutil
cpu: 12th Gen Intel(R) Core(TM) i5-12600K
BenchmarkGetDistance-16                 10000000                37.40 ns/op            0 B/op          0 allocs/op
BenchmarkGetDistanceHaversine-16        10000000                36.15 ns/op            0 B/op          0 allocs/op
BenchmarkGetBoundary-16                 10000000                18.12 ns/op           32 B/op          1 allocs/op
BenchmarkGeoHashEncode-16               10000000                87.57 ns/op           16 B/op          1 allocs/op
BenchmarkGeoHashDecode-16               10000000               184.6 ns/op            16 B/op          1 allocs/op
PASS
ok      github.com/tee-ck/go-geoutil    3.645s
```