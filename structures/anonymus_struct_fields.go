package main

import "fmt"

type Employee struct {
	string
	int
	float64
	designation string
}

func main() {

	employee := Employee{
		string:      "John Doe",
		int:         30,
		float64:     55000.50,
		designation: "Software Engineer",
	}
	fmt.Println("Name:", employee.string)
	fmt.Println("Age:", employee.int)
	fmt.Println("Salary:", employee.float64)
	fmt.Println("Designation:", employee.designation)

}
