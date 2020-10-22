package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start")
	// Timer represent a single event in the future
	// We can tell timer how long we wish to wait, and it
	// will return a channel which will be notified at time
	timer1 := time.NewTimer(2 * time.Second)

	// This blocks till timer fires on timer's channel C
	val := <-timer1.C
	fmt.Println(val)

	// If we just want to wait, we can use timer.Sleep
	// One of reasons to use timer is that it can be cancelled
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer2 fired!")
	}()
	// Stop timer2 before it fires
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer2 stopped")
	}

	// We wait to determine that timer2 really didn't fire and is
	// stopped
	time.Sleep(2 * time.Second)
}
