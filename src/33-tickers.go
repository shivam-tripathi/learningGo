package main

import (
	"fmt"
	"time"
)

// Tickers are used when we want to do something repeatedly
// at regular intervals

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Done")
				return
			case tick := <-ticker.C:
				fmt.Println("tick:", tick)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	// tickers can be stopped
	ticker.Stop()
	fmt.Println("Ticker Stopped", time.Now())
	done <- true
	time.Sleep(100 * time.Millisecond)
}
