package main

import "fmt"

func conditional() {
	theAnswer := 42
	var result string

	if theAnswer > 0 {
		result = "greater than zero"
	} else if theAnswer == 0 {
		result = "equal to zero"
	} else {
		result = "less than zero"
	}
	fmt.Println(result)

	// declaring the local variable within the if statement
	if local := "local"; local == "local" {
		fmt.Println("this is local variable")
	} else {
		fmt.Println("Not a local variable")
	}
}
