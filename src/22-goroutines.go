package main

import (
	"fmt"
	"time"
)

// goroutine is a lightweight thread of execution

func f(from string) {
	for i := 0; i < 10; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// Calling directly
	f("direct")

	// Calling concurrently using goroutine
	go f("goroutine")

	// goroutine can be used with anonymous function calls as well
	go func(msg string) {
		for i := 0; i < 20; i++ {
			fmt.Println(msg, ":", i)
		}
	}("going")

	// Wait for goroutines to complete
	time.Sleep(time.Second)
	fmt.Println("done")
}
