package main

import (
	"fmt"
	"os"
)

// Defer is used to ensure that a function call is performed
// later in a program's execution, usually for cleanup purposes

func createFile(p string) *os.File {
	fmt.Println("Creating file:", p)
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("Writing")
	fmt.Fprintf(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("Closing")
	err := f.Close()
	// Important check for errors even in deferred function
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	f := createFile("/tmp/file")
	defer closeFile(f)
	writeFile(f)
}
