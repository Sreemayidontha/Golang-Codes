package main

import "fmt"

func multiply(x, y int) int {
	return x * y
}
func main() {
	var a = 13
	var b = 7
	result := multiply(a, b)
	fmt.Printf("The multiplication of %d and %d is %d\n", a, b, result)
}
