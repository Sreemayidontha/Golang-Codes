package main

import "fmt"

var source :=[]int{10,20,30,40,50}
func main() {
	destination:=make([]int,len(source))
	for i:=;i<len(source);i++{
		destination[i]=source[i]
	}	
	fmt.Println("Source Slice:",source)
	fmt.Println("Before copy, destination:",destination)

}
