package main

import (
	"fmt"
	"os"
	"time"
)

// fail fast on errors we are not ready to handle gracefully

func main() {
	panic("some error occurred")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
