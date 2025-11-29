package main

import "fmt"

var source1 = []int{10, 20, 30, 40, 50}

func main() {
	desination := []int{}
	fmt.Println("Source Slice:", source1)
	fmt.Println("Before copy, destination:", desination)
	desination = append(desination, source1...)
	fmt.Println("After copy, destination:", desination)

}
