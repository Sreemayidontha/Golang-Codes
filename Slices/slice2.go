package main

import "fmt"

func main() {
	var a = []int{1, 2, 3, 4, 5, 6}
	b := a[:4]
	fmt.Println("The slice is a : ", a)
	fmt.Println("The slice is b : ", b)
	fmt.Printf("The capacity of b is %d\n", cap(b))
	b = append(b, 12, 13, 14, 15)
	fmt.Println("The slice is b : ", b)
	fmt.Printf("The capacity of b is %d\n", cap(b))
	//since the b capacity is greater than a, new array of double size of a is created
	// and values are copied

}
