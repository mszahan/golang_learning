package main

import (
	"fmt"
	"math"
)

func math_package() {
	f1, f2, f3 := 23.5, 65.1, 76.3
	sum := f1 + f2 + f3
	fmt.Println("Float sum:", sum)

	sum = math.Round(sum*100) / 100 // Round to one decimal places
	fmt.Println("Rounded sum:", sum)

	fmt.Println("The value of Pi is: ", math.Pi)

	circle_radius := 15.5
	circumference := 2 * math.Pi * circle_radius
	fmt.Printf("Circumference: %.2f\n", circumference) // %.2f formats the float to 2 decimal places
}
