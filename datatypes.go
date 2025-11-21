package main

import "fmt"

func main() {
	var a int = 10
	var b float32 = 20.5
	var c = "Golang"
	fmt.Println("Integer:", a)
	fmt.Println("Float:", b)
	fmt.Println("String:", c)
	fmt.Println(a-6, a+6)
	x := 176
	y := 23
	z := 0.00002345
	p := 123.45632
	fmt.Printf("The addition %d+%d =%d\n", x, y, x+y)
	fmt.Printf("The subtraction %d-%d=%d\n", x, y, x-y)
	fmt.Printf("The multiplication %d*%d=%d\n", x, y, x*y)
	fmt.Printf("The division %f/%f=%f\n", float64(x), float64(y), float64(x/y))
	fmt.Printf("The modulus %d%%%d=%d\n", x, y, x%y)
	fmt.Printf("The division %d/%d=%f\n", x, y, float64(x/y))
	fmt.Printf("The compact notation %g,%g\n", z, p)
	fmt.Printf("%.4g,%.7g", z, p)

}
