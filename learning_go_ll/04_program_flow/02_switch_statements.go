package main

import (
	"fmt"
	"time"
)

func switching() {
	weekday := time.Now().Weekday()
	fmt.Printf("Today is %v\n", weekday)

	day_number := int(weekday)
	fmt.Printf("The day as a number is %v\n", day_number)

	var result string
	switch day_number {
	case 1:
		result = "It's a Monday"
	case 2:
		result = "It's a Tuesday"
	case 3:
		result = "It's a Wednesday"
	case 4:
		result = "It's a Thursday"
	case 5:
		result = "It's a Friday"
	default:
		result = "It's a weekend"
	}

	fmt.Println(result)

	x := 42
	switch {
	case x < 0:
		result = "les than zero"
	case x == 0:
		result = "Equals zero"
	default:
		result = "Greater than zero"
	}

	fmt.Println(result)
}
