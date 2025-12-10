package main

import (
	"fmt"
)

func pf(c *int) {
	*c = 100

}
func pfd(d *int) {
	*d = 200

}
func main() {

	var a int = 10
	fmt.Println("Before:", a)
	var b *int = &a
	fmt.Println("Pointer Address:", b)
	fmt.Println("Value before :", *b)

	pf(b)
	fmt.Println("Value after :", a)
	var x int = 230
	fmt.Println("Before passing address of pointer :", x)
	pfd(&x)
	fmt.Println("Value after passing address of pointer :", x)
}
