package main

import "fmt"

// channels are pipes that connect concurrent go routines

func main() {
	messages := make(chan string)

	// Send value into channel
	go func() { messages <- "ping" }()

	// Receive value from channel
	msg := <-messages
	fmt.Println(msg)

	// By default channel sends and receives block until both sender
	// and receiver are ready. This allowed us to wait at the end of
	// program for message "ping" without any other synchronization
}
