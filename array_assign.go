package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	var b [3]int
	var c [3]int

	for i := 0; i < len(a); i++ {
		b[i] = a[i]
	}
	c = a

	fmt.Println("Array a:", a)
	fmt.Println("Array b (after element-wise assignment):", b)
	fmt.Println("Array c (after direct assignment):", c)

}
