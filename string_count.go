package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "GeeksforGeeks"
	str2 := "Welcome to GeeksforGeeks"
	str3 := "Data Science"

	res1 := strings.Count(str1, "e")
	res2 := strings.Count(str2, "Geeks")
	res3 := strings.Count(str3, "a")
	res4 := strings.Count("Hawaii islands are my favorite places to visit", "es")
	fmt.Println("Count of 'e' in str1: ", res1)
	fmt.Println("Count of 'Geeks' in str2: ", res2)
	fmt.Println("Count of 'a' in str3: ", res3)
	fmt.Println("Count of 'es' in str4: ", res4)
}
