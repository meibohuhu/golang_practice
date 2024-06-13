package main

import "fmt"

type Car struct {
	NumberOfTires int
	Luxury        bool
	BucketSeats   bool
	Make          string
	Model         string
	Year          int
}

func main() {

	// aggregate types (array, struct)
	// declare an array of strings
	var myStrings [3]string
	// add values to the array; remember that arrays start counting at 0
	myStrings[0] = "dog"
	myStrings[1] = "fish"
	myStrings[2] = "cat"
	// print out the first element of the array
	fmt.Println("The first item in the array is", myStrings[0])

	// reference types (pointers, slices, maps, functions, channels)

	// interface type

	// Car is the type for Cars
	// use Struct
	myCar := Car{
		NumberOfTires: 4,
		BucketSeats:   false,
		Make:          "Volvo",
		Model:         "XC90",
		Year:          2019,
	}
	fmt.Printf("My car is a %d %s %s", myCar.Year, myCar.Make, myCar.Model)
	fmt.Println()
}
