package main

import "fmt"

func calculate(a, b int) (int, int, int) {
	return a + b, a - b, a * b
}
func myfun(x, y int) (int, int) {
	square := x * x
	rectangle := x * y
	return square, rectangle
}
func main() {
	var1, var2, var3 := calculate(9, 6)
	fmt.Printf("Values returned are:%d\t%d\t%d", var1, var2, var3)

	area1, area2 := myfun(5, 7)
	fmt.Printf("\nSquare: %d\t Rectangle: %d", area1, area2)

}
