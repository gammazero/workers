package workers_test

import (
	"testing"
	"testing/synctest"

	"github.com/gammazero/workers"
)

func TestExample(t *testing.T) {
	do, done := workers.New(5)
	requests := []string{"alpha", "beta", "gamma", "delta", "epsilon"}

	for _, r := range requests {
		do <- func() {
			t.Log("Handling request:", r)
		}
	}
	close(do) // stop workers
	<-done    // wait for all workers to exit

	t.Log("All done")
}

func TestWorkers(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
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
				t.Errorf("Bad result for task %d", i)
			}
		}
		t.Log("All done")
	})
}

func TestBadNumberPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("worker.New(0) did not panic as expected")
		}
	}()

	workers.New(0)
}

func BenchmarkWorkers(b *testing.B) {
	do, done := workers.New(64)
	results := make([]bool, b.N)

	b.ResetTimer()
	for i := range b.N {
		do <- func() {
			results[i] = true
		}
	}
	close(do)
	<-done
	b.StopTimer()

	for i, ok := range results {
		if !ok {
			b.Errorf("Bad result for task %d", i)
		}
	}
}
