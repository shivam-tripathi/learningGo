package main

import (
	"fmt"
)

// A pipeline is nothing more than a series of things that take data in, perform an
// operation on it, and pass the data back out. We call each of these operations a
// stage of the pipeline
func main() {
	multiply := func(values []int, multiplier int) []int {
		multipliedValues := make([]int, len(values))
		for i, v := range values {
			multipliedValues[i] = multiplier * v
		}
		return multipliedValues
	}

	add := func(values []int, additive int) []int {
		addedValues := make([]int, len(values))
		for i, v := range values {
			addedValues[i] = additive + v
		}
		return addedValues
	}

	// Pipeline stages:
	// A stage consumes and returns the same type
	// A stage must be reified by the language so that it may be passed around.
	// Something like higher order functions (Go functions are reified)
	// Pipeline stages are like subsets of monads
	ints := []int{1, 2, 3, 4, 5}
	for _, v := range add(multiply(ints, 2), 1) {
		fmt.Println(v)
	}

	// batch processing -> processing chunks of data together
	// stream processing -> processing one by one

	fmt.Println("=====")
	generator := func(done <-chan interface{}, integers ...int) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for _, i := range integers {
				select {
				case <-done:
					return
				default:
					fmt.Println("Sending in", i)
					intStream <- i
				}
			}
			fmt.Println("Closing intstream")
		}()
		return intStream
	}

	multiplyChan := func(done <-chan interface{}, intStream <-chan int, multiplier int) <-chan int {
		multipliedStream := make(chan int)
		go func() {
			defer close(multipliedStream)
			for i := range intStream {
				select {
				case <-done:
					return
				case multipliedStream <- multiplier * i:
				}
			}
		}()
		return multipliedStream
	}

	addChan := func(done <-chan interface{}, intStream <-chan int, additive int) <-chan int {
		addedStream := make(chan int)
		go func() {
			defer close(addedStream)
			for {
				select {
				case <-done:
					return
				case i := <-intStream:
					addedStream <- additive + i
				}
			}
		}()
		return addedStream
	}

	done := make(chan interface{})
	defer close(done)

	intStream := generator(done, 1, 2, 3, 4, 5)
	pipeline := addChan(done, multiplyChan(done, intStream, 2), 1)
	for i := range pipeline {
		fmt.Println(i)
	}
}
