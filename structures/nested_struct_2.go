package main

import "fmt"

type Info struct {
	manufacturer string
	warranty     string
}
type Item struct {
	name    string
	price   float64
	device  string
	details Info
}

func main() {

	item := Item{name: "Laptop", price: 1200.50, device: "Electronics",
		details: Info{manufacturer: "TechCorp", warranty: "2 years"}}
	fmt.Printf("Item Name: %s", item.name)
	fmt.Printf("\nItem Price: %.2f", item.price)
	fmt.Printf("\nItem Manufacturer: %s", item.details.manufacturer)
}
