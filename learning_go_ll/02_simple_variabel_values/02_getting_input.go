package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func input() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println("You entered:", text)

	fmt.Print("Enter a number: ")
	text, _ = reader.ReadString('\n')
	f, err := strconv.ParseFloat(strings.TrimSpace(text), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("You entered the number:", f)
	}
}
