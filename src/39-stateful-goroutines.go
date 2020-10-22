package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// Go's idea of sharing memory by communicating where
// each piece of data is owned by exactly one goroutine

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	var readOps uint64
	var writeOps uint64

	reads := make(chan readOp)
	writes := make(chan writeOp)
	done := make(chan bool)
	result := make(chan map[int]int, 1)

	go func() {
		// State is owned by a single goroutine. This means data will not corrupted
		// by other goroutine, as a single goroutine executes synchronously, there
		// is not concurrent access.
		// Inorder to access the state, other goroutines will send messages to owning
		// goroutines and receive messages
		var state = make(map[int]int)
		for {
			select {
			case <-done:
				result <- state
				return
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// 100 readers which loop and try to read values
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read
				<-read.resp // wait till read is finished
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 10 writes which loop and try to write values
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool),
				}
				writes <- write
				<-write.resp // wait till write is finished
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOpsFinal:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOpsFinal:", writeOpsFinal)
	done <- true
	resultFinal := <-result
	fmt.Println("resultFinal:", resultFinal)
}
