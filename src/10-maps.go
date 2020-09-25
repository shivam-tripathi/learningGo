package main

import "fmt"

func main() {
	// maps are associated data types
	m := make(map[string]int)
	m["k1"] = 1
	m["k2"] = 2

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("m[k1]:", v1)

	fmt.Println("map len:", len(m))

	delete(m, "k1")

	_, isPresent := m["k1"]
	fmt.Println("isPresent:", isPresent)

	n := map[string]int{"foo": 1, "bar": 1}
	fmt.Println(n)
}
