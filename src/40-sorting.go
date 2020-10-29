package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Sorted strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Sorted ints:", ints)

	// We can also use sort package to check if slice is alredy sorted
	s := sort.IntsAreSorted(ints)
	fmt.Println("Ints are sorted?", s)
}
