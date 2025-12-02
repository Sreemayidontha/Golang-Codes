package main

import (
	"fmt"
	"sort"
)

func main() {
	intSlice := []int{42, 23, 16, 15, 8, 4}
	strslice := []string{"banana", "apple", "cherry", "date"}
	sort.Ints(intSlice)
	sort.Strings(strslice)
	fmt.Println("The ints sorte slices is ", intSlice)
	fmt.Println("The string sorted is: ", strslice)
	// Check again
	fmt.Println("sorted after sorting")
	fmt.Println("Ints:", sort.IntsAreSorted(intSlice))
	fmt.Println("Strings:", sort.StringsAreSorted(strslice))
}
