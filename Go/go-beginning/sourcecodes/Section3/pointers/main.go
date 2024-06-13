package main

import (
	"fmt"
)

// reference types (pointers, slices, maps, functions, channels)

// interface type

func main() {

	x := 10

	myFirstPointer := &x

	fmt.Println("x is", x)

	*myFirstPointer = 15

	fmt.Println("x is now", x)

	changeValueOfPointer(&x)

	fmt.Println("After function call, x is now", x)
}


func changeValueOfPointer(num *int) {
	*num = 25
}