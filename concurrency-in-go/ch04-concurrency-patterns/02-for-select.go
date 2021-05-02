package main

import (
	"fmt"
	"time"
)

func simple() {
	done := make(chan bool)
	defer close(done)
	stringstream := make(chan string)
	defer close(stringstream)
	go func() {
		for c := range stringstream {
			fmt.Println(c)
		}
	}()
	for _, c := range []string{"a", "b", "c", "d"} {
		select {
		case <-done:
			return
		case stringstream <- c:
		}
	}
}

// If the done channel is not closed, we will execute default case instead
// and get out of the select block
func infiniteLoop() {
	done := make(chan bool)
	defer close(done)
	go func() {
		time.Sleep(3 * time.Second)
		done <- true
	}()
	for {
		select {
		case <-done:
			return
		default:
			time.Sleep(1 * time.Second)
			// non preemptive work
		}
		// non preemptive work
		fmt.Print("Working \n")
	}
}

func main() {
	simple()
	infiniteLoop()
}
