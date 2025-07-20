package main

import "fmt"

func variables() {
	var price = 24
	var value int = 10

	new_price := 11

	if new_price == 10 {
		fmt.Println("the new price is 10")
	} else if new_price == 11 {
		fmt.Println("the new price is not 11")

	}

	fmt.Println(price)
	fmt.Println(value)
	fmt.Println(new_price)
}
