package main

import (
	"fmt"
	"strconv"
	"strings"
)

func calculate(value1 string, value2 string) float64 {
	// Your code goes here.
	f1, err := strconv.ParseFloat(strings.TrimSpace(value1), 64)

	if err != nil {
		fmt.Println("Error parsing first value:", err)
		panic("Value 1 must be a number")
	}

	f2, err := strconv.ParseFloat(strings.TrimSpace(value2), 64)
	if err != nil {
		fmt.Println("Error parsing first value:", err)
		panic("Value 2 must be a number")
	}

	result := f1 + f2

	return result
}
