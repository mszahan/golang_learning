package main

import (
	"fmt"
	"sort"
)

func check_slice() {
	// declare slice like array
	var colors = []string{"red", "green", "blue"}
	fmt.Println(colors)

	// more memory efficinet way of calling slice
	//make([]datatype, number_of_initial_items, capacity)
	var numbers = make([]int, 0, 3)
	numbers = append(numbers, 4, 7, 3)
	// you can append more elements more than capacity
	numbers = append(numbers, 1, 9)
	numbers = remove(numbers, 0)
	fmt.Println(numbers)

	// you can sort the slice
	sort.Strings(colors)
	fmt.Println(colors)
	sort.Ints(numbers)
	fmt.Println(numbers)

}

func remove(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}
