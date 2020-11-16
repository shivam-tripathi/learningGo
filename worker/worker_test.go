package worker

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestWorker(t *testing.T) {
	C := int64(0)
	var waitGroup sync.WaitGroup
	w := NewWorkerDefault("test")
	waitGroup.Add(1)
	w.Enqueue(func() {
		atomic.AddInt64(&C, 1000)
	}, &waitGroup)

	const N = 10
	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		w.Enqueue(func() {
			atomic.AddInt64(&C, 10)
		}, &waitGroup)
	}

	waitGroup.Wait()

	w.EnqueueSync(func() {
		atomic.AddInt64(&C, 999)
	})

	if C != (1000 + 10*10 + 999) {
		t.Fatal("Worker sum fail ", C)
		t.FailNow()
	}
}
