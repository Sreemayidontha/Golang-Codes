package main

import "fmt"

func main() {

	// Creating an array
	arr := [4]string{"Geeks", "for", "Geeks", "GFG"}

	// Creating slices from the given array
	var my_slice_1 = arr[1:2]
	my_slice_2 := arr[0:]
	my_slice_3 := arr[:2]
	my_slice_4 := arr[:]
	my_slice_5 := my_slice_4[1:3]

	// Display the result
	fmt.Println("My Array: ", arr)
	fmt.Println("My Slice 1: ", my_slice_1)
	fmt.Println("My Slice 2: ", my_slice_2)
	fmt.Println("My Slice 3: ", my_slice_3)
	fmt.Println("My Slice 4: ", my_slice_4)
	fmt.Println("My Slice 5: ", my_slice_5)
}
