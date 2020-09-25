package main

import "fmt"

// for is Go’s only looping construct.
func main() {
	// With a single condition
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// initial/condition/after loop
	for j:= 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// infinite loop (with break)
	for {
		fmt.Println("loop")
		break
	}

	// continue
	for n:= 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}
