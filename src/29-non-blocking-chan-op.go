package main

import (
	"fmt"
)

// Basic sends and receives on channels are blocking
// We can use select with a default clause to implement
// non blocking sends and non blocking multi way select

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// Non blocking receive
	// If a value is available on messages channel, it will
	// it, else it will immediately take up default case
	select {
	case msg := <-messages:
		fmt.Println("Received msg:", msg)
	default:
		fmt.Println("No messages received")
	}

	// Non blocking send
	// Here msg cannot be sent to messages channel because there is no
	// receiver and there is no buffer. Default case is selected.
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent messages", msg)
	default:
		fmt.Println("No messages sent")
	}

	// Multi way non blocking select
	select {
	case msg := <-messages:
		fmt.Println("Received message: ", msg)
	case sig := <-signals:
		fmt.Println("Received signal:", sig)
	default:
		fmt.Println("No activity")
	}
}
