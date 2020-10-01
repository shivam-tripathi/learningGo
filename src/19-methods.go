package main

import "fmt"

type rect struct {
	width, height int
}

// area method has a receiver type of rect
func (r *rect) area() int {
	return r.width * r.height
}

// methods can be defined on either pointer or value
func (r rect) perim() int {
	return 2 * (r.width + r.height)
}

func main() {
	r := rect{10, 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r

	// go handles conversion between pointer and values
	// Depending upon whether we want to copy or mutate struct on method call
	// we can choose between pointer receiver or value receiver
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
