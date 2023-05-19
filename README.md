# `xatomic`

[![GoDoc](https://godoc.org/github.com/go-ng/xatomic?status.svg)](https://pkg.go.dev/github.com/go-ng/xatomic?tab=doc)

## About

This package just provides with additional primitives to available in [`sync/atomic`](https://pkg.go.dev/sync/atomic).

Right now it provides with only two functions:

* [`LoadMap`](https://pkg.go.dev/github.com/go-ng/xatomic#LoadMap), which works similar to [`atomic.LoadUint64`](https://pkg.go.dev/sync/atomic#LoadUint64), but for a `map`, instead of `uint64`.
* [`StoreMap`](https://pkg.go.dev/github.com/go-ng/xatomic#StoreMap), which works similar to [`atomic.StoreUint64`](https://pkg.go.dev/sync/atomic#StoreUint64), but for a `map`, instead of `uint64`.

An example:

```go

import "github.com/go-ng/xatomic"


    ...
    var m map[string]any{}
    ...
    go func(){
        ...
        xatomic.StoreMap(m, myMap)
        ...
    }()
    go func(){
        ...
        myMap := xatomic.LoadMap(m)
        ...
    }()
```

It uses Go generics, so the type of the map is preserved after `StoreMap` & `LoadMap`.

## Performance

```plain
goos: linux
goarch: amd64
pkg: github.com/go-ng/xatomic
cpu: 11th Gen Intel(R) Core(TM) i7-11800H @ 2.30GHz
Benchmark/noop-16                     	1000000000	         0.2236 ns/op
Benchmark/Map/Store/atomic-16         	987858870	         6.108 ns/op
Benchmark/Map/Store/unatomic-16       	1000000000	         0.3973 ns/op
Benchmark/Map/Load/atomic-16          	1000000000	         0.4589 ns/op
Benchmark/Map/Load/unatomic-16        	1000000000	         0.4611 ns/op
PASS
ok  	github.com/go-ng/xatomic	8.364s
```
