package main

import "fmt"

func modify(num *int) {
	*num = 60
}
func main() {
	num := 20
	fmt.Printf("Before modification=%d\n", num)
	modify(&num)
	fmt.Printf("After modification=%d\n", num)
	fmt.Println("Address of num in main:", &num)
}
