// Package workers implements a simple fixed-size pool of goroutines for
// executing functions concurrently. The purpose is to limit the amount of
// concurrency to the number of goroutines in the pool.
package workers

import "sync"

// New creates a pool of goroutines that run the functions sent to them over
// the write channel. New returns two channels, a write-only "work" channel and
// a read-only "done" channel.
//
// The first channel is a write-only channel for sending task functions for the
// workers to run. When all workers are busy, the channel blocks. Close this
// channel To stop the workers.
//
// The second channel is a read-only channel that is closed when all workers
// have exited.
func New(numWorkers int) (chan<- func(), <-chan struct{}) {
	if numWorkers <= 0 {
		panic("number of workers must be greater than 0")
	}

	work := make(chan func())
	done := make(chan struct{})
	var wg sync.WaitGroup

	for range numWorkers {
		wg.Go(func() {
			for f := range work {
				f()
			}
		})
	}

	go func() {
		wg.Wait()
		close(done)
	}()

	return work, done
}
