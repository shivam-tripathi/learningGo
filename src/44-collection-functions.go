package main

import (
	"fmt"
	"strings"
)

func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}

	return vsf
}

func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {
	var strs = []string{"peach", "apple", "pear", "plum"}

	fmt.Println("Index of pear:", Index(strs, "pear"))
	fmt.Println("Includes guava:", Include(strs, "guava"))
	fmt.Println("Any one has prefix 'p':", Any(strs, func(str string) bool {
		return strings.HasPrefix(str, "p")
	}))
	fmt.Println("All have prefix 'p':", All(strs, func(str string) bool {
		return strings.HasPrefix(str, "p")
	}))
	fmt.Println("Contains 'e':", Filter(strs, func(str string) bool {
		return strings.Contains(str, "e")
	}))
	fmt.Println(Map(strs, strings.ToUpper))
	fmt.Println(strs)
}
