package main

import "fmt"

var ExportedVar = "Variable from File1"

func main() {
	fmt.Print("Hello, Sreemayi From file 1")
	fmt.Println(ExportedVar)
	fmt.Println(AnotherExportedVar)
}
