package main

import "fmt"

// aggregate types (array, struct)

// reference types (pointers, slices, maps, functions, channels)

// interface type

// Car is the type for Cars
type Car struct {
	NumberOfTires int
	Luxury        bool
	BucketSeats   bool
	Make          string
	Model         string
	Year          int
}

func main() {
	// declare an array of strings
	var myStrings [3]string

	// add values to the array; remember that arrays start counting at 0
	myStrings[0] = "dog"
	myStrings[1] = "fish"
	myStrings[2] = "cat"

	// print out the first element of the array
	fmt.Println("The first item in the array is", myStrings[0])

	// create a variable of type Car and populate it in one step
	myCar := Car{
		NumberOfTires: 4,
		Luxury:        true,
		BucketSeats:   true,
		Make:          "Volvo",
		Model:         "XC90",
		Year:          2019,
	}

	fmt.Printf("My car is a %d %s %s", myCar.Year, myCar.Make, myCar.Model)
	fmt.Println()
}
