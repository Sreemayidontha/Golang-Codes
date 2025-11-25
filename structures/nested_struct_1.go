package main

import "fmt"

type Student struct {
	Name    string
	class   string
	subject string
	Teacher Teacher
}
type Teacher struct {
	Name  string
	Email string
}

func main() {
	student := Student{
		Name:    "Bob",
		class:   "10th Grade",
		subject: "Mathematics",
		Teacher: Teacher{
			Name:  "Mr. Smith",
			Email: "smith@hotmail.com"},
	}
	fmt.Printf("The Student name is %s", student.Name)
	fmt.Printf("\nThe Student class is %s", student.class)
	fmt.Printf("\nThe Teacher name is %s", student.Teacher.Name)
}
