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

	p3 := new(Person)
	p3.name = "B"
	p3.age = 22
	fmt.Println(p3.name)
	fmt.Println(p3.age)
	fmt.Println(p1.age)
	fmt.Println(p2.age)

}
