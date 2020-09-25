package main

import (
	"fmt"
	"time"
)

func main() {
	// Simple switch
	i := 2
	fmt.Print("Write", i, "as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// Multi case switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// Condition less switch
	// case can contain expression besides constants
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// type switch
	// Variable assignment in condition
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'am a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI(1.2)
	whatAmI(int64(32))
	whatAmI(uint(22))

	// break exits the block
	v := 2
	switch v {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two before break")
		break
		fmt.Println("This won't print :: reachable code")
	default:
		fmt.Println("Default")
	}

	// fallthrough
	switch v {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two before break")
		fallthrough
	default:
		fmt.Println("Default") // this will also print
	}
}

