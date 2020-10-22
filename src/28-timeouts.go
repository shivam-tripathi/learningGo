package main

import (
	"fmt"
	"time"
)

// Timeouts are important for programs that connect to external resources
// or those which need to bound to execution time

func main() {
	// Timeout happens _before_ result
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result_1"
	}()
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout_1")
	}

	// Timeout happens _after_ result
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result_2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout_2")
	}
}
