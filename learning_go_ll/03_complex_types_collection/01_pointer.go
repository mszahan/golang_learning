package main

import "fmt"

func pointer() {
	anint := 42
	var p *int = &anint

	if p == nil {
		fmt.Println("pointer is nil")
	} else {
		fmt.Println("Value of p:", *p)
	}

	// changing value through pointer
	value1 := 32.13
	pointer1 := &value1
	*pointer1 = *pointer1 / 30
	fmt.Println("pointer: ", *pointer1)
	fmt.Println("value: ", value1)
}
