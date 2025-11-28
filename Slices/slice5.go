package main

import "fmt"

func main() {

	myslice := []string{"This", "is", "the", "tutorial", "of", "Go", "language"}
	for index, value := range myslice {
		fmt.Printf("The index is %d and the value is %s\n", index, value)
	}
	for index1, value1 := range myslice {
		fmt.Printf("The index is : %d  and the value is %v\n", index1+2, value1)

	}
}
