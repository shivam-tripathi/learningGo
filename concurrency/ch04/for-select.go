package main

import "fmt"

func simple() {
	done := make(chan bool)
	stringstream := make(chan string)
	for _, c := range []string{"a", "b", "c", "d"} {
		select {
		case <-done:
			return
		default:
			stringstream <- c
		}
	}

	for c := range stringstream {
		fmt.Println(c)
	}
}

func main() {
	simple()
}
