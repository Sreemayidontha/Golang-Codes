package main

import "fmt"

func main() {
	str1 := "GeeksforGeeks"
	str2 := "geeksForgeeks"
	str3 := "GeeksforGeeks"
	str4 := "Sreemayi"
	result1 := str1 == str2
	result2 := str1 == str3
	fmt.Printf("The name is %s", str1)
	fmt.Printf("\nThe comparison between %s and %s is %t", str1, str2, result1)
	fmt.Printf("\nThe comparison between %s and %s is %t", str1, str3, result2)
	fmt.Printf("\nThe type of str1 is %T", str1) //T for type, t for true/false
	fmt.Printf("\nThe lenghth of the str3 is %d", len(str3))
	fmt.Printf("\nThe combination of the whole string is %s", str1+" "+str4)
}
