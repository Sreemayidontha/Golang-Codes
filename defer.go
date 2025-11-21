package main

import "fmt"

func multiply(a, b int) {
	res := a * b
	fmt.Println("The multiplication is:", res)
}
func show() {
	fmt.Println("This is show function")
}

func main() {

	fmt.Println("Defer functions")
	multiply(4, 5)

	defer multiply(56, 89)
	show()
}
