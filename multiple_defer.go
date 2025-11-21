package main

import "fmt"

func multiply(a, b int) {
	res := a * b
	fmt.Println("The multiplication is:", res)
}

func main() {

	fmt.Println("Start")
	defer fmt.Println("End")
	defer multiply(34, 78)
	defer fmt.Println("Multiple defer functions are LIFO type")
	defer multiply(5, 3)

}
