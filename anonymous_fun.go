package main

import "fmt"

func GFG(a string) string {
	return "Geeks"

}
func main() {
	func() {
		fmt.Println("Anonymous Function Called")
	}()
	value := func(msg string) {
		fmt.Println("Hello " + msg)
	}
	value("Sreemayi")
	message := func(p, q string) string {
		return p + q + GFG("Geeks")
	}
	last := func() {
		fmt.Println(message("Geeks", "For"))
	}
	last()

}
