package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	var b *[3]int = &a
	var c *[3]int
	c = b

	fmt.Println("Array a:", a)
	fmt.Println("Pointer b points to array:", *b)
	fmt.Println("Pointer c points to array:", b)
	fmt.Println("Pointer c points to array:", *c)
	fmt.Println("Pointer c points to array:", c)
}
