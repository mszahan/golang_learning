package main

import (
	"fmt"
	"time"
)

func date_and_time() {

	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("Go language launched at", t)

	n := time.Now()
	fmt.Println("Current time:", n)
	fmt.Print(n.Format(time.ANSIC) + "\n")

	tomorrow := n.AddDate(0, 0, 1)
	fmt.Println(tomorrow.Format(time.ANSIC))

	format := "Sunday, Jan 2, 2006"
	fmt.Println("Formatted date: ", tomorrow.Format(format))
}
