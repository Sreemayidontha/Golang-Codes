package main

import (
	"fmt"
)

func main() {
	var a int = 58
	var b *int = &a
	fmt.Println("Value of a:", a)
	fmt.Println("Address of a:", &a)
	fmt.Println("Value of b (Address of a):", b)
	fmt.Println("Value pointed to by b:", *b)
	c := 1103
	var d = &c
	fmt.Println("Value of c:", c)
	fmt.Println("Address of c:", &c)
	fmt.Println("Value of d (Address of c):", d)
	fmt.Println("Value pointed to by d:", *d)
	e := &a
	fmt.Println("Value of e (Address of a):", e)
	fmt.Println("Value pointed to by e:", *e)
}
