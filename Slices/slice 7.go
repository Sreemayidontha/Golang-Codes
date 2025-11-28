package main

import (
	"fmt"
	"sort"
)

func main() {
	s1 := []int{12, 34, 56}
	var s2 []int
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
	s3 := [][]int{{1, 2, 3}, {4, 5, 6}}
	s4 := [][]int{[]int{1, 2}, []int{3, 4}}
	fmt.Println(s3)
	fmt.Println(s4)
	//sort operations names are same as the data types
	a := []int{5, 2, 6, 3, 1, 4}
	b := []string{"banana", "apple", "mango", "grape"}
	fmt.Println("Unsorted a:", a)
	fmt.Println("Unsorted b:", b)
	sort.Ints(a)
	sort.Strings(b)
	fmt.Println("Sorted a:", a)
	fmt.Println("Sorted b:", b)

}
