package main

import "fmt"

func main() {
	var ptr1 [7]*int
	var ptr2 [5]*string
	var ptr3 [8]*float64

	fmt.Println("Capacity of ptr1: ", cap(ptr1))
	fmt.Println("Capacity of ptr2: ", cap(ptr2))
	fmt.Println("Capacity of ptr3: ", cap(ptr3))
	fmt.Println("Length of ptr1: ", len(ptr1))

}
