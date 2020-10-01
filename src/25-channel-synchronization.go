package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("Working...")
	time.Sleep(time.Second * 2)
	fmt.Println("Done")

	done <- true
}

func main() {
	done := make(chan bool)
	go worker(done)
	<-done // Block until worker is finished
	fmt.Println("Exiting")
}
