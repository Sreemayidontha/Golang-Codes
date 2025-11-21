package main

import "fmt"

func main() {
	var intVar int = 42
	var floatVar float64 = 3.14
	b := float64(intVar) // Converting int to float64
	c := int(floatVar)   // Converting float64 to int
	var p int32 = 28
	var q int = p
	var r int64 = 100
	var add = p + q + r // This will cause a compilation error due to mismatched types
	fmt.Println("Integer Variable:", intVar)
	fmt.Println("Float Variable:", floatVar)
	fmt.Println("Converted Float from Int:", b)
	fmt.Println("Converted Int from Float:", c)
	fmt.Println("Sum of p, q, r:", add)
	fmt.Printf("add:%f", floatVar+float64(c))
}
