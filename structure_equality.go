package main

import (
	"fmt"
	"reflect"
)

type Author struct {
	Name string
	Book string
	Age  int
}

func main() {
	a1 := Author{"John Doe", "Go Programming", 45}
	a2 := Author{"John Doe", "Go Programming", 45}
	a3 := Author{"Jane Smith", "Python Programming", 30}

	if a1 == a2 {
		fmt.Println("A1 and A2 author are same")
	} else {
		fmt.Println("A1 and A2 author are different")
	}
	fmt.Println("Is A1 and A3 authors same: ", reflect.DeepEqual(a1, a3))
	fmt.Println("Is A1 and A2 authors same: ", reflect.DeepEqual(a1, a2))
	a1.Age = 50
	fmt.Println("After modifying age of A1:", a1)

}
