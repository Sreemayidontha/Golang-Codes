// Anonumous Go routines

package main

import (
	"fmt"
	"time"
)

func main() {
	go func(s string) {
		for i := 0; i < 5; i++ {

			fmt.Println(s)
			time.Sleep(500 * time.Millisecond)
		}
	}("Hello from Anonymous Goroutine!")
	go func(s string) {
		for i := 0; i < 5; i++ {

			fmt.Println(s)
			time.Sleep(30 * time.Millisecond)
		}

	}("Another Anonymous Goroutine!")
	time.Sleep(1 * time.Second)
	fmt.Println("Main function complete.")
	//Goroutines in Go let functions run concurrently, using less memory
	// than traditional threads. Every Go program starts with a main Goroutine,
	// and if it exits, all others stop.
}
