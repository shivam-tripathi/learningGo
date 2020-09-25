package main

import "fmt"

/*
- Arrays are value types, i.e. on reassignment they are deep copied.
- Length is part of array's type. [3]int and [4]int are completely different types.
*/

func main() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	// declare and initialize
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// compose types to build multi-dimensional array
	var twoD [2][3]int
	for i:=0; i < 2; i++ {
		for j:=0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}
