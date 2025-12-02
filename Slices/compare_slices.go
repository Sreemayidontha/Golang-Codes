package main

import (
	"bytes"
	"fmt"
)

func main() {

	slice1 := []byte{'G', 'o', 'l', 'a', 'n', 'g'}
	slice2 := []byte{'G', 'o', 'l', 'a', 'n', 'g'}
	slices := []byte{'G', 'o', 'l', 'a', 'n', 'g', '!'}

	fmt.Println(bytes.Equal(slice1, slice2))
	fmt.Println(bytes.Equal(slice1, slices))
	fmt.Println(bytes.Compare(slice1, slice2))

}
