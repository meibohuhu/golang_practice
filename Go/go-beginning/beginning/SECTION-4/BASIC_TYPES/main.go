package main

import "log"

// basic types (numbers, strings, booleans)
var myInt int
var myUint uint
var myFloat float32
var myFloat64 float64

// aggregate types (array, struct)

// reference types (pointers, slices, maps, functions, channels)

// interface type

func main() {
	myInt = 10
	myUint = 20

	myFloat = 10.1
	myFloat64 = 100.1

	log.Println(myInt, myUint, myFloat, myFloat64)

	myString := "Trevor"

	log.Println(myString)

	myString = "John"

	var myBool = true
	// myBool = 0 // will be wrong
	myBool = false

	log.Println(myBool)

}
