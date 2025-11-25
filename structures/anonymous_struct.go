package main

import "fmt"

type Student struct {
	personaldetails struct {
		name string
		age  int
	}
	GPA float32
}

func main() {

	student := Student{
		personaldetails: struct {
			name string
			age  int
		}{
			name: "Alice",
			age:  20,
		},
		GPA: 3.8,
	}
	fmt.Println("Name:", student.personaldetails.name)
	fmt.Println("Age:", student.personaldetails.age)
	fmt.Println("GPA:", student.GPA)
}
