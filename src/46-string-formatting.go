package main

import (
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

func main() {
	p := Point{1, 2}

	// Print instance of struct
	fmt.Printf("%v\n", p)

	// Include struct field names
	fmt.Printf("%+v\n", p)

	// Go syntax representational value
	fmt.Printf("%#v\n", p)

	// Print type of value
	fmt.Printf("%T\n", p)

	// Formatting boolean
	fmt.Printf("%t\n", true)

	// Formatting number: base 10
	fmt.Printf("%d\n", 23)

	// Formatting number: base 2
	fmt.Printf("%b\n", 19)

	// Character corresponding to given integer
	fmt.Printf("%c\n", 48)

	// Formatting number: hex encoding
	fmt.Printf("%x\n", 456)

	// Formatting floats: base 10
	fmt.Printf("%f\n", 34.23)

	// Formatting floats: Formatting number in scientific notation
	fmt.Printf("%e\n", 124.23)
	fmt.Printf("%E\n", 124.23)

	// Formatting string: simple
	fmt.Printf("%s\n", "\"string\"")

	// Formatting string: double quotes as in string (ingore escape)
	fmt.Printf("%q\n", "\"string\"")

	// Formatting string: hex (base 16)
	fmt.Printf("%x\n", "string")

	// Representation of pointer
	fmt.Printf("%p\n", &p)

	// Representation of number with width: By default right aligned
	fmt.Printf("%3d\n", 23)
	fmt.Printf("%3d\n", 123)

	// Representation of floats: width.precision (width = nondecimal + decimal)
	fmt.Printf("%8.3f\n", 123.33)

	// Use - flag to left align
	fmt.Printf("[%-3d]\n", 23)
	fmt.Printf("[%-6s]\n", "foo")

	// Use Sprintf to format and return string without using it anywhere else
	s := fmt.Sprintf("%b", 31)
	fmt.Printf("%s\n", s)

	// Format and print to io.Writers other than os.stdout using Fprintf
	fmt.Fprintf(os.Stderr, "Stderr via fmt: an %s\n", "error")
}
