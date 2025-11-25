package main

import "fmt"

type details struct {
	course string
	grade  string
}
type School struct {
	name string
	age  int
	details
}

func main() {
	s := School{
		name: "Alice",
		age:  20,
		details: details{
			course: "Computer Science",
			grade:  "A",
		},
	}

	fmt.Println("Name:", s.name)
	fmt.Println("Age:", s.age)
	fmt.Println("Course:", s.course) // Accessing promoted field directly
	fmt.Println("Grade:", s.grade)   // Accessing promoted field directly
}
