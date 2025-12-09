package main

import (
	"fmt"
	"strings"
)

func main() {

	str1 := "Welcome to the online portal of GeeksforGeeks"
	str2 := "My dog name is Dollar"
	str3 := "I like to play Ludo"

	fmt.Println("String 1: ", str1)
	fmt.Println("String 2: ", str2)
	fmt.Println("String 3: ", str3)

	res1 := strings.IndexAny(str1, "G")
	res2 := strings.IndexAny(str2, "do")
	res3 := strings.IndexAny(str3, "lqxa")
	res4 := strings.IndexAny("GeeksforGeeks, geeks", "uywq")

	res5 := strings.Index(str1, "Geeks")
	res6 := strings.Index(str2, "do")
	res7 := strings.Index(str3, "chess")
	res8 := strings.Index("GeeksforGeeks, geeks", "ks")

	// Displaying the result
	fmt.Println("\nIndex values:")
	fmt.Println("Result 1: ", res1)
	fmt.Println("Result 2: ", res2)
	fmt.Println("Result 3: ", res3)
	fmt.Println("Result 4: ", res4)
	fmt.Println("Result 5: ", res5)
	fmt.Println("Result 6: ", res6)
	fmt.Println("Result 7: ", res7)
	fmt.Println("Result 8: ", res8)

}
