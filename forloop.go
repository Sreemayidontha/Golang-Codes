package main

import "fmt"

func main() {
	var name = "Sreemayi Dontha"
	for i, j := range name {
		fmt.Printf("The index is %d and the character is %c\n", i, j)
	}
	for i := 0; i <= len(name)-1; i++ {
		fmt.Printf("The index is %d and the character is %c\n", i, name[i])
	}
	k := 0
	for k < 6 {
		fmt.Println("I belong somewhere beautiful and best like always wanted")
		k++
	}

}
