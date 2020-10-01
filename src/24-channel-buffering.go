package main

import "fmt"

func main() {
	// Buffered channels receive a limited number of values without corresponding receiver

	// Buffering upto 2 values
	messages := make(chan string, 2)

	messages <- "hello"
	messages <- "world"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
