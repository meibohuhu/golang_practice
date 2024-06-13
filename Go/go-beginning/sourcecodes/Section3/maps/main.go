package main

import "fmt"

// reference types (pointers, slices, maps, functions, channels)

// interface type

func main() {
	// maps are reference types, so they are always passed by reference. You don't need a pointer.
	// maps are keys and values, and you must use the make function when creating
	intMap := make(map[string]int)

	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3
	intMap["four"] = 4
	intMap["five"] = 5

	// range through the map and print out key and value
	for key, value := range intMap {
		fmt.Println(key, value)
	}

	// delete a value from a map by key
	delete(intMap, "four")

	// check to see if an item exists in a map by key
	el, ok := intMap["four"]
	if ok {
		fmt.Println(el, "is in map")
	} else {
		fmt.Println(el, "is in not map")
	}

	// update a value in the map
	intMap["two"] = 4
}
