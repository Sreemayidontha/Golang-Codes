package main

import "fmt"

func init() {
	fmt.Println("Init function called - Initialization logic here")
}

func init() {
	fmt.Println("Another init() function, which doesn't conflict with the first one and doesnt depend on main package initialization ")
}

func main() {
	fmt.Println("Main function called - Main program logic here")
}
