# workers
> simple worker pool

[![GoDoc](https://pkg.go.dev/badge/github.com/gammazero/workers)](https://pkg.go.dev/github.com/gammazero/workers)
[![Build Status](https://github.com/gammazero/workers/actions/workflows/go.yml/badge.svg)](https://github.com/gammazero/workers/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gammazero/workers/branch/main/graph/badge.svg)](https://codecov.io/gh/gammazero/workers)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/gammazero/workers/blob/main/LICENSE)

The `workers` package implements a simple fixed-size pool of goroutines for executing functions concurrently. The purpose is to limit concurrency to the number of goroutines in the pool.

When all workers are busy, the channel for submitting tasks blocks. If you are looking for a worker pool that never blocks when submitting tasks, see [workerpool](https://github.com/gammazero/workerpool).

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
	requests := []string{"alpha", "beta", "gamma", "delta", "epsilon"}

	for _, r := range requests {
		do <- func() {
			fmt.Println("Handling request:", r)
		}
	}
	close(do) // stop workers
	<-done    // wait for all workers to exit

	fmt.Println("All done")
}
```
