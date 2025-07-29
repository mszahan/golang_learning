package main

import "fmt"

func test_array() {
	var colors [3]string
	colors[0] = "red"
	colors[1] = "green"
	colors[2] = "blue"

	fmt.Println(colors)
	fmt.Println("first color", colors[0])

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	fmt.Println("Number of colors", len(colors))
	fmt.Println("Number of numbers", len(numbers))
}
