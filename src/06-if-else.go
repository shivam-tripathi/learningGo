package main

import "fmt"

func main() {
	// Parantheses are not required, but brackets are
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// A statement can precede conditionals, available
	// to all branches
	if num := 9; num < 0 {
		fmt.Println(num, "num is negative")
	} else if num < 10 {
		fmt.Println(num, "num has 1 digit and is positive")
	} else {
		fmt.Println(num, "num has multiple digits and is positive")
	}
}

