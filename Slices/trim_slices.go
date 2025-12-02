package main

import (
	"bytes"
	"fmt"
)

func main() {
	slice1 := []byte("  Hello, Gophers!  ")
	slice2 := []byte{'!', '!', 'G', 'e', 'e', 'k', 's', 'f',
		'o', 'r', 'G', 'e', 'e', 'k', 's', '#', '#'}
	slice3 := []byte{'*', '*', 'A', 'p', 'p', 'l', 'e', '^', '^'}
	slice4 := bytes.Trim([]byte("****Welcome to GeeksforGeeks****"), "*")
	res2 := bytes.Trim([]byte("!!!!Learning how to trim a slice of bytes@@@@"), "!@")
	res3 := bytes.Trim([]byte("^^Geek&&"), "$")

	r1 := bytes.TrimSpace(slice1)
	r2 := bytes.Trim(slice2, "#!")
	r3 := bytes.Trim(slice3, "^*")
	fmt.Printf("Original Slice1: %s\n", slice1)
	fmt.Printf("Trimmed Slice1: %s\n", r1)
	fmt.Printf("Original Slice2: %s\n", slice2)
	fmt.Printf("Trimmed Slice2: %s\n", r2)
	fmt.Printf("Original Slice3: %s\n", slice3)
	fmt.Printf("Trimmed Slice3: %s\n", r3)
	fmt.Printf("Trimmed Slice4: %s\n", slice4)
	fmt.Printf("Trimmed Result2: %s\n", res2)
	fmt.Printf("Trimmed Result3: %s\n", res3)
}
