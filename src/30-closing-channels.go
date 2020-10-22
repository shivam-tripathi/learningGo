package main

import (
	"fmt"
)

// Closing channel can be useful to use as a signal that
// no more values will be sent to it

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("Received job:", j)
			} else {
				fmt.Println("All jobs received.")
				done <- true
				return
			}
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Println("Sending job: ", i)
		jobs <- i
	}
	close(jobs)
	fmt.Println("Sent all jobs")

	// Await worker using synchronization
	<-done
}
