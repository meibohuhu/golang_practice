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

	fmt.Println(animals)

	for i, x := range animals {
		fmt.Println(i, x)
	}

	fmt.Println("Element 0 is", animals[0])

	fmt.Println("First two elements are", animals[0:2])

	fmt.Println("The slice is", len(animals), "elements long")

	fmt.Println("Is it sorted?", sort.StringsAreSorted(animals))

	sort.Strings(animals)

	fmt.Println("Is it sorted now?", sort.StringsAreSorted(animals))
	fmt.Println(animals)

	animals = DeleteFromSlice(animals, 1)
	fmt.Println(animals)
}

func DeleteFromSlice(a []string, i int) []string {
	a[i] = a[len(a)-1] // first, copy the last element in the slice to index i
	a[len(a)-1] = ""   // next, erase the last element by setting to an empty string
	a = a[:len(a)-1]   // finally, truncate the slice by getting all elements except for the last one
	return a           // return the truncated slice
}
