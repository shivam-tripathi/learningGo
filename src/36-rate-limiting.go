package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	// Rate limited request (every 200 milliseconds)
	for req := range requests {
		<-limiter
		fmt.Println("request:", req, time.Now())
	}

	// allow bursts of request by using buffering
	burstyLimiter := make(chan time.Time, 3)
	// Burst of 3 requests allowed
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	// every 200ms we'll try to add a new request, upto a limit
	// of 3
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()
	// Simulate burstyRequests
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	// Simulate incoming requests - first 3 will be immediately handled
	// Remaining will come after every 200ms
	fmt.Println()
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request:", req, time.Now())
	}
}
