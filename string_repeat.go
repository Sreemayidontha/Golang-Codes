package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "Sreemayi Dontha"
	s2 := strings.Repeat(s1, 3)
	fmt.Println(s2)
	s3 := s2 + " Welcome to GoLang"
	fmt.Println(s3)
	s4 := "Data Expert" + strings.Repeat(s1, 5)
	fmt.Println(s4)
	str1 := " Hello"
	res1 := " "
	i := 0
	for i < 3 {
		res1 += str1
		i++

	}
	fmt.Println(res1)

}
