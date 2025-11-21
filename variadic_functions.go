package main

import "fmt"

func sum(num ...int) int {

	total := 0
	for i := range num {
		total += num[i]
	}
	return total
}
func greet(message string, nums ...int) {
	fmt.Println(message, len(message), nums, len(nums))
}
func main() {
	fmt.Println("The sum of numbers are :", sum(1, 2, 3, 4, 5))
	fmt.Println("The sum of numbers are :", sum(10, 20))
	greet("Passing string", 1, 2, 4, 5, 6)
	greet("No numbers passed")

}
