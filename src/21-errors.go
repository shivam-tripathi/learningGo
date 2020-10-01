package main

import (
	"errors"
	"fmt"
)

// Errors are explicitly returned
// By default, they are last return value and have type `error`

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Cannot work with 42")
	}
	// nil in position of error indicates there is no error
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

// Custom error can be created by implementing the Error method in the error interface
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "Cannot work with 42"}
	}
	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 passed:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 passed:", r)
		}
	}

	// If we want to use data in a custom error, we will have to get the error as
	// an instance of the custom error type via type assertion
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
