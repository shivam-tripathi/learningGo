package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("p[a-z]+ch", "peach")
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")
	peachPunch := "peach punch"
	peachPunchPinch := "peach punch pinch"

	fmt.Println(r.MatchString(peachPunch))

	fmt.Println(r.FindString(peachPunch))

	// Returns empty slice if not found
	fmt.Println(r.FindStringIndex(peachPunch))

	// Whole patterns matches + submatches in them
	fmt.Println(r.FindStringSubmatch(peachPunch))

	// Returns indexes for full match + submatches
	fmt.Println(r.FindStringSubmatchIndex(peachPunch))

	// Find more than one matches. Second argument limits number of matches
	// If negative, all matches are returned
	fmt.Println(r.FindAllString(peachPunchPinch, -1))
	fmt.Println(r.FindAllString(peachPunchPinch, 2))

	// Can be run bytes as well, remove 'String' from function names
	fmt.Println(r.Match([]byte(peachPunch)))

	// While using regex as global variable, use must compile. It panics instead
	// of returning error
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	// Replace occurences of strings
	fmt.Println(r.ReplaceAllString(peachPunch, "<fruit>"))

	// Replace all func allows to transform all matched strings
	in := []byte("a peach in punch with pinch")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(in), "->", string(out))
}
