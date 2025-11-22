package main

import "fmt"

type Student struct {
	Name   string
	Age    int
	Grade  string
	Rollno int
}
type Cars struct {
	Brand, model, color string
	WeightKg            float64
}

func main() {
	s1 := Student{"Alice", 20, "A", 101}
	s2 := Student{"Bob", 22, "B", 102}

	fmt.Println("Student 1:", s1)
	fmt.Println("Student 2:", s2)
	c1 := Cars{Brand: "BMW", model: "M series", color: "Blue"}
	c2 := Cars{"Audi", "A8", "Black", 1800.5}

	fmt.Println("Car 1:", c1)
	fmt.Println("Car 2:", c2)
	fmt.Println("The model of the car 2 is : ", c2.model)
	s3 := &Student{"Charlie", 23, "C", 103}
	fmt.Println("Student 3:", *s3)
}
