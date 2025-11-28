package main

import "fmt"

func main() {
	s1 := []int{12, 34, 56}
	var s2 []int
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
	s3 := [][]int{{1, 2, 3}, {4, 5, 6}}
	s4 := [][]int{[]int{1, 2}, []int{3, 4}}
	fmt.Println(s3)
	fmt.Println(s4)

}
