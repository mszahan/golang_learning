package main

import "fmt"

func check_structs() {
	poodle := Dog{"Poodle", 34}
	fmt.Println(poodle)
	// print both key and value
	fmt.Printf("%+v\n", poodle)
	// getting each value with . notation
	poodle_bread := poodle.Breed
	poodle_weight := poodle.Weight
	fmt.Println(poodle_bread)
	fmt.Println(poodle_weight)

	// change the value with . notation
	poodle.Weight = 45
	fmt.Println(poodle)
}

type Dog struct {
	Breed  string
	Weight int
}
