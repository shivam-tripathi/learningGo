package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Primary mechanism for managing state in go is communication
// over channels. Another way is using sync/atomic package.

func main() {
	var ops uint64
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1) // Atomically increment the counter
				// ops += 1 // This would have led to different final value
			}
			wg.Done()
		}()
	}

	wg.Wait()

	// Safe to read ops now, as no go rountine is accessing it
	// We can use atomic.LoadUint64 to access atomically a variable as well
	// Using race detector, we would error if we used ops++ instead above
	fmt.Println("ops:", ops)
}
