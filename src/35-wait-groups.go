package main

import (
	"fmt"
	"sync"
	"time"
)

// To wait for multiple goroutines, we can use wait group

// WaitGroup must be passed as a pointer
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d started \n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d stopped\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
}
