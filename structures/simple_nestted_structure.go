package main

import "fmt"

type Author struct {
	Name string
	Book string
	Age  int
}

type HR struct {
	details Author
}

func main() {
	result := HR{details: Author{Name: "Alice", Book: "Learning Go", Age: 29}}
	fmt.Println("HR Details:", result)
}
