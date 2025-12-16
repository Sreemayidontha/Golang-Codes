package main

import "fmt"

func display(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

func main() {

	go display("Goroutine 1")
	go display("Goroutine 2")
	display("The main go")
}
