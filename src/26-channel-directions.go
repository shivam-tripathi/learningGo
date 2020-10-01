package main

import "fmt"

// when using channels in function parameters, we can specify if
// channel is to be used for sending only or for only receiving as well
// This helps in type safety
func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	fmt.Println("Received from ping, sending to pong", msg)
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	go ping(pings, "pinged message")
	go pong(pings, pongs)

	fmt.Println(<-pongs)
}
