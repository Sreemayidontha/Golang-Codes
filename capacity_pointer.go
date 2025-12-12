package main

import "fmt"

func main() {
	arr := [8]int{200, 300, 400,
		500, 600, 700, 100, 200}
	fmt.Println("Capacity of arr: ", cap(arr))

	var b [5]*int
	fmt.Println("Capacity of b: ", cap(b))
	var x int
	for x := 0; x < len(b); x++ {
		b[x] = &arr[x]
	}
	for x = 0; x < len(b); x++ {
		fmt.Printf("Value of b[%d] =%d\n", x, *b[x])
	}
	fmt.Println("Capacity of b :", cap(b))

}
