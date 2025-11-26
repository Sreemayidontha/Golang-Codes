package main

import "fmt"

func main() {
	a := [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(a)
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			fmt.Printf("a[%d][%d]=%d ", i, j, a[i][j])
		}
	}
	var b [1][2]string
	b[0][0] = "Hello"
	b[0][1] = "World"

	fmt.Println(b)
}
