package main

import (
	"fmt"

	"github.com/Sreemayidontha/Golang/golang"
)

var ExportedVar = "Variable from File1"

func main() {
	fmt.Print("Hello, Sreemayi From file 1")
	fmt.Println(ExportedVar)
	fmt.Println(golang.AnotherExportedVar)
}
