package main

import "fmt"

func main() {
	var name = "Sreemayi Dontha"
	i := 0
	for i < len(name) {
	Label1:
		fmt.Printf("The name is %s\n", name)
		if name[i] == ' ' {
			goto Label1
		}
		if i == 12 {
			break
		}
		if i == 5 {
			i = i + 4
			continue
		}
		fmt.Printf("value is: %d\n", i)
		i++
	}

}
