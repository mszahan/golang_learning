package main

import "fmt"

func loops() {
	fmt.Println("loops")

	colors := []string{"Red", "Green", "Blue"}

	// traditional loop
	for i := 0; i < len(colors); i++ {
		println(colors[i])
	}

	// range in loop
	for i := range colors {
		println(colors[i])
	}

	// key value type accessing range
	for index, color := range colors {
		println(index, color)
	}

	states := make(map[string]string)
	states["WA"] = "Washington"
	states["OR"] = "Oregon"
	states["Ca"] = "California"

	for key, value := range states {
		println(key, value)
	}

	// for loop as while loop in other language
	value := 0
	sum := 0

	for value < 5 {
		sum += value
		fmt.Printf("Value: %v\n", value)
		fmt.Printf("Sum: %v\n", sum)
		value++
	}

	// for loop with break an continue type
	sum = 1
	for sum < 1000 {
		sum += sum
		if sum > 200 {
			goto theEnd
		}
	}
theEnd:
	println("end of program")
	fmt.Printf("Sum: %v\n", sum)
}
