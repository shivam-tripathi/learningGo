package main

import "fmt"

// structs are typed collection of fields
// They're useful for grouping data together to form records

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 44
	return &p
}

func main() {
	fmt.Println(person{name: "Tony", age: 32})
	fmt.Println(person{"Kate", 28})
	fmt.Println(person{name: "McGee"})
	fmt.Println(&person{name: "Mallard", age: 55})
	fmt.Println(newPerson("Gibbs"))

	s := person{name: "Abby", age: 24}
	fmt.Println(s.age)

	// we can use dots with struct pointers
	// the pointers are automatically dereferenced.
	sp := &s
	fmt.Println(sp.name)

	sp.age = 30
	fmt.Println(sp.age)
}
