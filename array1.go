package main

import "fmt"

func main() {

	a := [4]int{1, 2, 3, 4}
	fmt.Println(a)
	b := []string{"Go", "Python", "Java"}

	for i := 0; i < len(b); i++ {
		fmt.Println(b[i])
	}
}
