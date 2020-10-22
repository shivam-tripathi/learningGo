package main

import (
	"fmt"
)

// We can use for range to iterate over channel

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// Because channel is closed, iteration terminates
	for elem := range queue {
		fmt.Println(elem)
	}

	// It's possible to close a non empty channel but still have
	// the remaining values received
}
