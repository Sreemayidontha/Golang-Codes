// Go program to illustrate how to
// check the equality of the slices

package main

import (
	"bytes"
	"fmt"
)

func main() {

	// Creating and initializing
	// slices of bytes
	// Using shorthand declaration
	slice_1 := []byte{'A', 'N', 'M',
		'A', 'P', 'A', 'A', 'W'}

	slice_2 := []byte{'g', 'e', 'e', 'k', 's'}

	slice_3 := []byte{'A', 'N', 'M', 'A',
		'P', 'A', 'A', 'W'}

	// Checking the equality of the slices
	// Using Equal function
	res1 := bytes.Equal(slice_1, slice_2)
	res2 := bytes.Equal(slice_1, slice_3)
	res3 := bytes.Equal(slice_2, slice_3)
	res4 := bytes.Equal([]byte("GeeksforGeeks"),
		[]byte("GeeksforGeeks"))
	res5 := bytes.Equal([]byte("Geeks"), []byte("GFG"))
	res6 := bytes.Equal(slice_1, []byte("P"))

	// Displaying results
	fmt.Println("Result 1:", res1)
	fmt.Println("Result 2:", res2)
	fmt.Println("Result 3:", res3)
	fmt.Println("Result 4:", res4)
	fmt.Println("Result 5:", res5)
	fmt.Println("Result 6:", res6)

}
