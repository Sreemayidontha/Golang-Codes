package main

import "fmt"

func po_array(b *[5]int) {
	(*b)[3] = 345
}
func main() {
	var a [5]int = [5]int{10, 20, 30, 40, 50}
	fmt.Println("Before:", a)

	po_array(&a)
	fmt.Println("After:", a)

	//Array pointer are recommended in Go as they are difficult to read. Use slices instead.
}
