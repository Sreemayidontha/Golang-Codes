package main

import "fmt"

func calavg(a [6]int) int {
	var sum int
	b := len(a)
	for _, v := range a {
		sum += v
	}
	return sum / b
}
func main() {
	var a = [6]int{67, 59, 29, 35, 4, 34}

	avg := calavg(a)
	fmt.Println("Average is:", avg)
}
