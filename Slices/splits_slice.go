package main

import (
	"bytes"
	"fmt"
)

func main() {

	slice1 := []byte("a,b,b")
	slice2 := []byte("a;b;c")
	slice3 := []byte("abc")
	res1 := bytes.Split(slice1, []byte(","))
	res2 := bytes.Split(slice2, []byte(";"))
	res3 := bytes.Split(slice3, []byte{})
	fmt.Printf("Original Slice1: %s\n", slice1)
	for _, part := range res1 {
		fmt.Println(string(part))
	}
	fmt.Printf("Original Slice2: %s\n", slice2)
	for _, part := range res2 {
		fmt.Println(string(part))
	}
	fmt.Printf("Original Slice3: %s\n", slice3)
	for _, part := range res3 {
		fmt.Println(string(part))
	}

}
