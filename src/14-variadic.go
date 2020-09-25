package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(11, 22)
	sum(11, 22, 33)

	nums := []int{11, 22, 33, 44}
	sum(nums...)
}
