package main

import "fmt"

func main() {
	var a = []int{1, 2, 3, 4, 5, 6}
	b := a[:4]
	fmt.Println("The slice is : ", a)
	fmt.Printf("The capacity of a is %d\n", cap(a))
	fmt.Printf("Length of a is %d\n", len(a))
	fmt.Println("The slice b is : ", b)
	fmt.Printf("The capacity of b is %d\n", cap(b))
	fmt.Printf("Length of b is %d\n", len(b))
}
