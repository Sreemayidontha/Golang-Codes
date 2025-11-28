package main

import "fmt"

func main() {
	var a = make([]int, 4, 7)
	b := make([]int, 0)
	c := make([]int, 5)
	fmt.Printf("The slice a:%d \t capacity:%d \t length:%d\n", a, cap(a), len(a))
	fmt.Printf("The slice b:%d \t capacity:%d \t length:%d\n", b, cap(b), len(b))
	fmt.Printf("The slice c:%d \t capacity:%d \t length:%d\n", c, cap(c), len(c))
}
