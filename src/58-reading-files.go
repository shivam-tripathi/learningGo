package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Slurping entire content of file into memory
	dat, err := ioutil.ReadFile("./README.md")
	check(err)
	fmt.Println(string(dat))

	// We often want more control how and what part of file is read
	// For this first we open file
	f, err := os.Open("./README.md")
	check(err)

	// Read some random bytes from file
	b1 := make([]byte, 14) // allow upto five bytes
	n1, err := f.Read(b1)  // n1 is number of bytes read
	check(err)
	fmt.Printf("number of bytes read: %d @  0: Read [%s]\n", n1, string(b1))

	// Seek to a new location in the file and read from there
	// 0 -> relative to origin, 1 -> relative to current offset, 2 -> relative to end
	o2, err := f.Seek(218, 0)
	check(err)
	b2 := make([]byte, 8)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("number of bytes read: %d @ %d: ", n2, o2)
	fmt.Printf("Read [%s]\n", string(b2[:n2]))

	o3, err := f.Seek(218, 0)
	check(err)
	b3 := make([]byte, 8)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Println(fmt.Sprintf("number of bytes read: %d @ %d: Read [%s]", n3, o3, string(b3)))

	r4 := bufio.NewReader(f)
	o4, err := f.Seek(218, 0)
	check(err)
	b4, err := r4.Peek(8)
	check(err)
	fmt.Printf("number of bytes read: 8 @ %d: Read [%s]\n", o4, string(b4))

	// rewind
	_, err = f.Seek(0, 0)
	check(err)
}
