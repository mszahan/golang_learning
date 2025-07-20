package main

import "fmt"

func simple_math() {
	i1, i2, i3 := 10, 20, 30

	int_sum := i1 + i2 + i3
	fmt.Println("Sum of integers:", int_sum)

	float1, float2 := 10.5, 20.6
	float_sum := float1 + float2
	fmt.Println("Sum of floats:", float_sum)

	sum_float_int := float64(int_sum) + float_sum
	fmt.Println("Sum of float and int:", sum_float_int)

	int_float_div := float64(i1) / float1
	fmt.Println("Division of int by float:", int_float_div)
}
