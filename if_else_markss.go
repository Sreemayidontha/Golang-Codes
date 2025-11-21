package main

import "fmt"

func main() {
	var mark1 float64 = 85.5
	var mark2 float64 = 45.0
	var mark3 float64 = 50.0
	var average float64 = (mark1 + mark2 + mark3) / 3
	fmt.Printf("Average Mark: %.2f\n", average)

	if average >= 90 {
		fmt.Println("Grade A")
		if mark1 > 95 && mark2 > 95 && mark3 > 95 {
			fmt.Println("Outstanding!")
		} else {
			fmt.Println("Great job!")
		}
	} else if average >= 80 {
		fmt.Println("Grade B")
	} else if average >= 70 {
		fmt.Println("Grade C")
	} else if average >= 60 {
		fmt.Println("Grade D")
	} else {
		if average < 40 {
			fmt.Println("Fail — Needs serious improvement")
		} else {
			fmt.Println("Fail — Almost passed")
		}
	}

}
