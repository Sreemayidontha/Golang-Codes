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
	time.Sleep(1 * time.Second) // Allow Goroutine to finish
	fmt.Println("Main function complete.")
}
