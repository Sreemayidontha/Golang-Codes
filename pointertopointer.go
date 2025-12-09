package main

import (
	"fmt"
)

func main() {
	var a int = 42
	var b *int = &a
	var c **int = &b
	fmt.Println("Value of a:", a)
	fmt.Println("Address of a:", &a)
	fmt.Println("Value of b (Address of a):", b)
	fmt.Println("Value pointed to by b:", *b)
	fmt.Println("Value of c (Address of b):", c)
	fmt.Println("Value pointed to by c (Address of a):", *c)
	fmt.Println("Value pointed to by the pointer pointed to by c (Value of a):", **c)
	fmt.Println(c)

}
