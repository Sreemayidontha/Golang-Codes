package main

import "fmt"

func main() {

	val1 := 2345
	val2 := 567

	var p1 *int
	p1 = &val1
	p2 := &val2
	p3 := &val1

	res1 := &p1 == &p2
	fmt.Println("Is p1 pointer is equal to p2 pointer: ", res1)
	fmt.Println("Address of p1 pointer: ", &p1)
	fmt.Println("Address of p2 pointer: ", &p2)
	fmt.Println(p1)
	fmt.Println(p2)

	res2 := p1 == p2
	fmt.Println("Is p1 pointer is equal to p2 pointer: ", res2)

	res3 := p1 == p3
	fmt.Println("Is p1 pointer is equal to p3 pointer: ", res3)

	res4 := p2 == p3
	fmt.Println("Is p2 pointer is equal to p3 pointer: ", res4)

	res5 := &p3 == &p1
	fmt.Println("Is p3 pointer is equal to p1 pointer: ", res5)

	res6 := &p1 != &p2
	fmt.Println("Is p1 pointer not equal to p2 pointer: ", res6)

	res7 := p1 != p2
	fmt.Println("Is p1 pointer not equal to p2 pointer: ", res7)

	res8 := p1 != p3
	fmt.Println("Is p1 pointer not equal to p3 pointer: ", res8)

	res9 := p2 != p3
	fmt.Println("Is p2 pointer not equal to p3 pointer: ", res9)

	res10 := &p3 != &p1
	fmt.Println("Is p3 pointer not equal to p1 pointer: ", res10)

}
