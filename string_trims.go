package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "@@Hello, Geeks!!"
	r1 := strings.Trim(s, "@!")
	fmt.Println(r1)
	r2 := strings.TrimLeft(s, "@!")
	fmt.Println(r2)
	r3 := strings.TrimRight(s, "Geeks!")
	fmt.Println(r3)
	s2 := "   Hello, Geeks   "
	r4 := strings.TrimSpace(s2)
	fmt.Println(r4)

}
