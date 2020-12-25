package main

import (
	"bytes"
	"fmt"
	"sync"
)

// Confinement ensures information only available from one concurrent process
// Ad hoc confinement is achieved by a covention
// Eg: Here we are only accessing data using loopData, even though it is
// possible to do so using data slice
func adHocConfinement() {
	data := []int{1, 2, 3, 4}
	loopData := func(hndlData chan<- int) {
		defer close(hndlData)
		for i := range data {
			hndlData <- data[i]
		}
	}

	hndlData := make(chan int)
	go loopData(hndlData)

	for num := range hndlData {
		fmt.Println(num)
	}
}

// Lexical confinement involves using lexical scope to expose only expose only the
// correct data and concurrency primitives for multiple concurrent processes to use
func lexicalConfinement() {
	// Channel is defined within the lexical scope of chanOwner, we return only read
	// only channel. This prevents other goroutines from writing to this channel
	chanOwner := func() <-chan int {
		results := make(chan int, 5)
		go func() {
			defer close(results)
			for i := 0; i < 5; i++ {
				results <- i
			}
		}()
		return results
	}

	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Println(result)
		}
		fmt.Println("Done")
	}

	results := chanOwner()
	consumer(results)
}

// printData doesn't close around the data slice, it needs slice of byte to operate on
// This confines printData to slice of byte we pass to it, and it cannot edit the data
// byte array. As a result we can write syncronous code within printData and no critical
// sections are required.
func lexicalConfinement2() {
	printData := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()
		var buf bytes.Buffer
		for _, b := range data {
			fmt.Fprintf(&buf, "%c", b)
		}
		fmt.Println(buf.String())
	}
	var wg sync.WaitGroup
	wg.Add(2)
	data := []byte("golang")
	go printData(&wg, data[:3])
	go printData(&wg, data[3:])
	wg.Wait()
}

func main() {
	adHocConfinement()
	lexicalConfinement()
	lexicalConfinement2()
}
