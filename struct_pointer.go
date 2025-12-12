package main

import "fmt"

// Defining a struct
type Person struct {
	name string
	age  int
}

func main() {

	p1 := Person{name: "A", age: 25}
	p2 := &p1
	fmt.Println(p1.age)
	fmt.Println(p2.age)
	p2.age = 30
	fmt.Println(p1.age) // Output: 30

}
