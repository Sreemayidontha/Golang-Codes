package main

import "fmt"

var source = []int{10, 20, 30, 40, 50}

func main() {
	destination := make([]int, len(source))

	slice := copy(destination, source)
	fmt.Println("Source Slice:", source)
	fmt.Println("Before copy, destination:", destination)
	fmt.Println("Number of elements copied:", slice)
}
