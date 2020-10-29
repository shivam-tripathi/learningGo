package main

import (
	"fmt"
	"sort"
)

// To sort by custom function, implement sort.Interface
// Functions required: Len, Less, Swap
// Len and Swap will usually be same across, Less will be
// different, and effect sort order

type byLength []string

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
