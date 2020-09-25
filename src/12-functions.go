package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(33, 22)
	fmt.Println("33+22 =", res)

	res = plusPlus(44, 33, 22)
	fmt.Println("44+33+22 =", res)
}
