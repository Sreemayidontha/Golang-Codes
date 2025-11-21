package main

import "fmt"

func multiply(a, b int) (int, int) {
	return a * b, a + b
}

func main() {

	var mul, _ = multiply(5, 6)
	fmt.Println("The multiplication value is:", mul)
}
