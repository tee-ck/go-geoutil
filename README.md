# go-geoutil

## A Go library for geolocation utilities.

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