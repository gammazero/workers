# workers
> simple worker pool

[![GoDoc](https://pkg.go.dev/badge/github.com/gammazero/workers)](https://pkg.go.dev/github.com/gammazero/workers)
[![Build Status](https://github.com/gammazero/workers/actions/workflows/go.yml/badge.svg)](https://github.com/gammazero/workers/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gammazero/workers/branch/main/graph/badge.svg)](https://codecov.io/gh/gammazero/workers)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/gammazero/workers/blob/main/LICENSE)

The `workers` package implements a simple fixed-size pool of goroutines for executing functions concurrently. The purpose is to limit the amount of concurrency to the number of goroutines in the pool.

## Installation

```
$ go get github.com/gammazero/workers
```

## Example

```go
package main

import (
	"fmt"
	"github.com/gammazero/workers"
)

func main() {
	do, done := workers.New(5)

	const ntasks = 10
	results := make([]bool, ntasks)

	for i := range ntasks {
		do <- func() {
			t.Log("hello", i)
			results[i] = true
		}
	}
	close(do) // stop workers
	<-done    // wait for all workers to exit

	for i, ok := range results {
		if !ok {
			fmt.Println("Bad result for task %d", i)
		}
	}
	fmt.Println("All done")
}
```
