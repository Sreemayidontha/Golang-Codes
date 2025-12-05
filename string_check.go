package main

import (
	"fmt"
	"strings"
)

func main() {

	str1 := "Welcome to Geeks for Geeks"
	str2 := "Here! we learn about go strings"
	r1 := strings.Contains(str1, "to")
	fmt.Println("Contains:", r1)
	r2 := strings.HasPrefix(str1, "Welcome")
	fmt.Println("Has Prefix:", r2)
	r3 := strings.HasSuffix(str2, "strings")
	fmt.Println("Has Suffix:", r3)
	r4 := strings.Index(str1, "Geeks")
	fmt.Println("Index of substring:", r4)
	r5 := strings.LastIndex(str1, "Geeks")
	fmt.Println("Last Index of substring:", r5)
	r6 := strings.Count(str1, "e")
	fmt.Println("Count of substring:", r6)
	r7 := strings.EqualFold("goLang", "GOLANG")
	fmt.Println("Equal Fold:", r7)
	r8 := strings.ContainsAny(str1, "u & e")
	fmt.Println("Contains Any:", r8)
}
