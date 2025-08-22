package main

import (
	"fmt"
	"sort"
)

func check_maps() {
	// make(map[type_of_key]type_of_value)
	states := make(map[string]string)
	states["WA"] = "Washington"
	states["OR"] = "Oregon"
	states["CA"] = "California"
	fmt.Println(states)

	// read item from the map
	california := states["CA"]
	fmt.Println(california)

	// delete item from the map
	delete(states, "OR")
	fmt.Println(states)

	states["NY"] = "New York"

	// access data through for loop
	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}

	// getting the values in the alphabetical order of the key
	// seems too much work though
	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println("\nSorted")

	for i := range keys {
		fmt.Println(states[keys[i]])
	}

}
