package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "Hello"
	s2 := "Geeks"
	s3 := "Hello"

	// Equality and inequality
	fmt.Println("s1 == s2:", s1 == s2) // false
	fmt.Println("s1 == s3:", s1 == s3) // true
	fmt.Println("s1 != s2:", s1 != s2) // true

	// Lexicographical comparison
	fmt.Println("s1 < s2:", s1 < s2)                                 // true
	fmt.Println("s1 > s2:", s1 > s2)                                 // false
	fmt.Println("s1 <= s3:", s1 <= s3)                               // true
	fmt.Println("s1 >= s3:", s1 >= s3)                               // true
	fmt.Println("strings.Compare(s1, s2):", strings.Compare(s1, s2)) // -1
	fmt.Println("strings.Compare(s1, s3):", strings.Compare(s1, s3)) // 0
	fmt.Println("strings.Compare(s2, s1):", strings.Compare(s2, s1)) // 1
}
