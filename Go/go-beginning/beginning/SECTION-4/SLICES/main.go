package main

import (
	"fmt"
	"sort"
)

// reference types (pointers, slices, maps, functions, channels)

// interface type

func main() {
	var animals []string
	animals = append(animals, "dog")
	animals = append(animals, "fish")
	animals = append(animals, "cat")
	animals = append(animals, "horse")

	for i, x := range animals {
		fmt.Println(i, x)
	}

	fmt.Println("First two elements are", animals[0:2])
	fmt.Println("The slice is", len(animals), "elements long")

	fmt.Println("Is it sorted?", sort.StringsAreSorted(animals)) // false
	sort.Strings(animals)
	fmt.Println("Is it sorted?", sort.StringsAreSorted(animals)) // true

	DeleteFromSlice(animals, 1) // delete index, but mess up the order

	fmt.Println(animals)
	fmt.Println("Is it sorted?", sort.StringsAreSorted(animals)) // false
	sort.Strings(animals)
	fmt.Println("Is it sorted?", sort.StringsAreSorted(animals)) // true
}

func DeleteFromSlice(slices []string, index int) {
	slices[index] = slices[len(slices)-1]
	slices[len(slices)-1] = ""
	slices = slices[0 : len(slices)-1]
	// return slices
}
