package main

import "fmt"

func main() {
	a := [3]int{10, 20, 30}
	b := [...]string{"apple", "banana", "cherry"}
	c := [3]int{10, 20, 30}
	e := [3]string{"apple", "banana", "orange"}

	fmt.Println(a == c)
	fmt.Println(len(a))
	fmt.Println(b == e)
	fmt.Println(len(c))

}
