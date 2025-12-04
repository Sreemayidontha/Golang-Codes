package main

import "fmt"

func main() {

	My_value_1 := "Welcome to GeeksforGeeks"
	var My_value_2 string
	My_value_2 = "GeeksforGeeks"
	// Displaying strings
	fmt.Println("String 1: ", My_value_1)
	fmt.Println("String 2: ", My_value_2)
	My_value_3 := `Hello!GeeksforGeeks`

	// Adding escape character
	// in raw literals
	My_value_4 := `Hello!\nGeeksforGeeks`
	fmt.Println("String 3: ", My_value_3)
	fmt.Println("String 4: ", My_value_4)
	for index, s := range "GeeksForGeeKs" {

		fmt.Printf("The index number of %c is %d\n", s, index)
	}
	for c := 0; c < len(My_value_3); c++ {
		fmt.Printf("\nCharacter = %c Bytes = %x", My_value_3[c], My_value_3[c])
	}

}
