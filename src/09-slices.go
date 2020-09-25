package main

import "fmt"

func main() {
	// slices are typed by the elements they contain
	// not by the length. They are reference types unlike
	// arrays which are value types
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// Can be accessed via index
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// len return length
	fmt.Println("len:", len(s))

	// append returns a new slice (non-mutating)
	// append adds an extra element
	s2 := append(s, "d")
	s = append(s, "d", "e", "f")
	fmt.Println("apd:", s, s2)

	// slices can be copied
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// slice operation [low:high]
	l := s[2:5] // 2, 3, 4 (2<=x<5)
	fmt.Println(l)

	l = s[:5] // 0, 1, 2, 3, 4 (x<5)
	fmt.Println(l)

	l = s[2:] // 2, 3, 4, 5 (2<=x)
	fmt.Println(l)

	// slice declare and initialize
	// (similar to array but without size)
	t := []string{"g","h","i"}
	fmt.Println("dcl:", t)

	// slices can be composed to make 2D slices
	// length of inner slice can vary, unlike arrays
	twoD := make([][]int, 3)
	for i:= 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}
