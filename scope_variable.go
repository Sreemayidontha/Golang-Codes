package main

import "fmt"

var globalVar = "I am a global variable"

func main() {
	fmt.Println(globalVar) // Accessing global variable
	localVar := "I am a local variable"
	fmt.Println(localVar) // Accessing local variable
}
func display() {
	// fmt.Println(localVar) // This would cause an error: undefined localVar
	fmt.Println(globalVar) // Accessing global variable inside another function
}
